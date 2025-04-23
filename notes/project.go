package notes

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const DefaultProjectFile = "project.json"

func LoadProject(projectPath string) (*Project, error) {
	project, err := LoadJSON[Project](projectPath)
	if err != nil {
		return nil, err
	}
	project.loadedPath = projectPath
	return &project, nil
}

type Project struct {
	Name string `json:"name"`

	LogsPath     string `json:"logPath"`
	TasksPath    string `json:"taskPath"`
	MeetingsPath string `json:"meetingPath"`
	TopicsPath   string `json:"topicPath"`
	ImagesPath   string `json:"imagePath"`
	BuildsPath   string `json:"buildPath"`

	Logs     []*Log     `json:"logs"`
	Tasks    []*Task    `json:"tasks"`
	Meetings []*Meeting `json:"meetings"`
	Topics   []*Topic   `json:"topics"`
	Images   []*Image   `json:"images"`
	Tags     []*Tag     `json:"tags"`

	// The path to this project on disk
	loadedPath string
}

func (p Project) SetupFS(parentFolder string) error {
	projectFolderName := removeNonAlphanumeric(p.Name)
	folder := filepath.Join(parentFolder, projectFolderName)
	err := os.Mkdir(folder, os.ModeDir)
	if err != nil {
		return fmt.Errorf("unable to create project folder %s: %w", folder, err)
	}

	return SaveJSON(filepath.Join(folder, DefaultProjectFile), p)
}

func (p *Project) ListTasks(out io.Writer) error {
	for i, task := range p.Tasks {
		status := "Created"
		statusTime := task.Created
		if len(task.History) > 0 {
			item := task.History[len(task.History)-1]
			status = string(item.Status)
			statusTime = item.Time
		}

		_, err := fmt.Fprintf(out, "[%d] %-10s %s - %s\n", i+1, status, statusTime.Format("2006-01-02 15:04"), task.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) ListTodoTasks(out io.Writer) error {
	for i, task := range p.Tasks {
		status := "Created"
		statusTime := task.Created
		if len(task.History) > 0 {
			item := task.History[len(task.History)-1]
			status = string(item.Status)
			statusTime = item.Time
		}

		notFinished := status == "Created" || status == TaskStatus_Started
		if !notFinished {
			continue
		}

		_, err := fmt.Fprintf(out, "[%d] %-10s %s - %s\n", i+1, status, statusTime.Format("2006-01-02 15:04"), task.Name)
		if err != nil {
			return err
		}
	}
	return nil
}

type ProjectCompileOptions struct {
	UseMarkdownItIncludeExtension bool
	Save                          bool
}

func roundUpToDays(d time.Duration) int {
	return int(math.Round(float64(d) / float64(24*time.Hour)))
}

func (p *Project) TaskGantt(out io.Writer) {
	now := time.Now()

	fmt.Fprintln(out, "gantt")
	fmt.Fprintln(out, "    title Task Work History")
	fmt.Fprintln(out, "    dateFormat  YYYY-MM-DD")
	fmt.Fprintln(out, "    excludes weekends")

	for i, t := range p.Tasks {
		status := ""
		completed := ""
		var latest *TaskStatusChange
		if len(t.History) > 0 {
			latest = t.History[len(t.History)-1]
			switch latest.Status {
			case TaskStatus_Started:
				status = "active,"

			case TaskStatus_Completed:
				status = "done,"
				completed = latest.Time.Format("2006-01-02")

			case TaskStatus_Abandoned:
				continue
			}
		}

		started := ""
		for _, history := range t.History {
			if history.Status == TaskStatus_Started {
				started = history.Time.Format("2006-01-02")
				break
			}
		}

		if started == "" {
			continue
		}

		if status == "active," {
			completed = fmt.Sprintf("%dd", roundUpToDays(now.Sub(latest.Time)))
		}

		fmt.Fprintf(out, "    %s :%s desc%d, %s, %s\n", t.DisplayName(), status, i, started, completed)
	}
}

func (p *Project) compile_toc(out io.Writer) {
	fmt.Fprint(out, "## Table of Contents\n\n")

	fmt.Fprint(out, "* [Tasks](#tasks)\n")
	for i, task := range p.Tasks {
		fmt.Fprintf(out, "    * [%s](#tasks-%d)\n", task.DisplayName(), i)
	}

	fmt.Fprint(out, "* [Meetings](#meetings)\n")
	for i, meeting := range p.Meetings {
		fmt.Fprintf(out, "    * [%s](#meetings-%d)\n", meeting.Created.Format("2006-01-02 15:04"), i)
	}

	fmt.Fprint(out, "* [Topics](#topics)\n")
	for i, topic := range p.Topics {
		fmt.Fprintf(out, "    * [%s](#topics-%d)\n", topic.Name, i)
	}

	fmt.Fprint(out, "* [Logs](#logs)\n")
	for i, log := range p.Logs {
		fmt.Fprintf(out, "    * [%s](#logs-%d)\n", log.Path, i)
	}

	fmt.Fprint(out, "\n")
}

func (p *Project) Compile(out io.Writer, options ProjectCompileOptions) error {

	writer := bufio.NewWriter(out)
	if options.Save {
		folderName := filepath.Join(filepath.Dir(p.loadedPath), p.BuildsPath, time.Now().Format("2006-01-02 15 04"))
		err := os.MkdirAll(folderName, os.ModeDir)
		if err != nil {
			return err
		}

		f, err := os.Create(filepath.Join(folderName, "Project.md"))
		if err != nil {
			return err
		}
		defer f.Close()
		writer = bufio.NewWriter(f)
	}

	useIncludeDirective := options.UseMarkdownItIncludeExtension

	fmt.Fprintf(writer, "# %s\n\n", p.Name)

	fmt.Fprint(writer, "Tags: ")
	for i, tag := range p.Tags {
		fmt.Fprint(writer, tag.Name)
		if i < len(p.Tags)-1 {
			fmt.Fprint(writer, ", ")
		}
	}
	fmt.Fprint(writer, "\n\n")

	p.compile_toc(writer)

	fmt.Fprint(writer, "## <a id=\"tasks\">Tasks</a>\n\n")

	fmt.Fprint(writer, "```mermaid\n")
	p.TaskGantt(writer)
	fmt.Fprint(writer, "```\n\n")

	for i, task := range p.Tasks {
		fmt.Fprintf(writer, "### <a id=\"tasks-%d\">%s</a>\n\n", i, task.DisplayName())
		fmt.Fprintf(writer, "*Created: %s*\n\n", task.Created)

		for _, item := range task.History {
			fmt.Fprintf(writer, "* %s: %s", item.Status, item.Time)

			if item.Reason != "" {
				fmt.Fprintf(writer, " - %s", item.Reason)
			}
			fmt.Fprint(writer, "\n")
		}

		if len(task.History) > 0 {
			fmt.Fprint(writer, "\n")
		}

		descriptionPath := filepath.Join(filepath.Dir(p.loadedPath), p.TasksPath, task.Path, taskFileName)

		if !useIncludeDirective {
			description, err := os.ReadFile(descriptionPath)
			if err != nil {
				return err
			}
			writer.Write(description)
		} else {
			fmt.Fprintf(writer, ":[Task](%s)\n", descriptionPath)
		}

		fmt.Fprint(writer, "\n")
	}

	fmt.Fprint(writer, "## <a id=\"meetings\">Meetings</a>\n\n")
	for i, meeting := range p.Meetings {

		fmt.Fprintf(writer, "### <a id=\"meetings-%d\">%s</a>\n\n", i, meeting.Created.Format("2006-01-02 15:04"))
		fmt.Fprintf(writer, "*Created: %s*\n\n", meeting.Created)

		meetingPath := filepath.Join(filepath.Dir(p.loadedPath), p.MeetingsPath, meeting.Path, meetingFileName)
		if !useIncludeDirective {
			logData, err := os.ReadFile(meetingPath)
			if err != nil {
				return err
			}
			writer.Write(logData)
		} else {
			fmt.Fprintf(writer, ":[Meeting](%s)\n", meetingPath)
		}
		fmt.Fprint(writer, "\n")
	}

	fmt.Fprint(writer, "## <a id=\"topics\">Topics</a>\n\n")
	for i, topic := range p.Topics {

		fmt.Fprintf(writer, "### <a id=\"topics-%d\">%s</a>\n\n", i, topic.Name)
		fmt.Fprintf(writer, "*Created: %s*\n\n", topic.Created)

		topicPath := filepath.Join(filepath.Dir(p.loadedPath), p.TopicsPath, topic.Path, topicFileName)

		if !useIncludeDirective {
			logData, err := os.ReadFile(topicPath)
			if err != nil {
				return err
			}
			writer.Write(logData)
		} else {
			fmt.Fprintf(writer, ":[Topic](%s)\n", topicPath)
		}
		fmt.Fprint(writer, "\n")
	}

	fmt.Fprint(writer, "## <a id=\"logs\">Logs</a>\n\n")
	for i := len(p.Logs) - 1; i >= 0; i-- {
		log := p.Logs[i]

		fmt.Fprintf(writer, "### <a id=\"logs-%d\">%s</a>\n\n", i, log.Path)
		fmt.Fprintf(writer, "*Created: %s*\n\n", log.Created)

		logPath := filepath.Join(filepath.Dir(p.loadedPath), p.LogsPath, log.Path, logFileName)
		if !useIncludeDirective {
			logData, err := os.ReadFile(logPath)
			if err != nil {
				return err
			}
			writer.Write(logData)
		} else {
			fmt.Fprintf(writer, ":[Log](%s)\n", logPath)
		}
		fmt.Fprint(writer, "\n")
	}
	return writer.Flush()
}

func (p *Project) NewLog(tags []string) error {
	t := time.Now()
	log := &Log{
		Created: t,
		Path:    t.Format("2006-01-02"),
		Tags:    sanitizeTags(tags),
	}

	err := log.initiailzeMarkdown(filepath.Join(filepath.Dir(p.loadedPath), p.LogsPath))
	if err != nil {
		return fmt.Errorf("error creating log for project %s: %w", p.Name, err)
	}

	p.Logs = append(p.Logs, log)
	return nil
}

func (p *Project) OpenLog() error {
	t := time.Now()
	logPath := t.Format("2006-01-02")
	for _, log := range p.Logs {
		if log.Path == logPath {
			openURL(filepath.Join(filepath.Dir(p.loadedPath), p.LogsPath, logPath, logFileName))
			return nil
		}
	}

	return p.NewLog(nil)
}

func (p *Project) NewMeeting() error {
	t := time.Now()
	meeting := &Meeting{
		Created: t,
		Path:    t.Format("2006-01-02 15 04"),
	}

	err := meeting.initiailzeMarkdown(filepath.Join(filepath.Dir(p.loadedPath), p.MeetingsPath))
	if err != nil {
		return fmt.Errorf("error creating meeting for project %s: %w", p.Name, err)
	}

	p.Meetings = append(p.Meetings, meeting)
	return nil
}

func (p *Project) AddImage(originalImagePath string) error {
	t := time.Now()
	image := &Image{
		Entry: Entry{
			Created: t,
			Path:    fmt.Sprintf("%d%s", len(p.Images), filepath.Ext(originalImagePath)),
		},
		OriginalPath: originalImagePath,
	}

	imagePath := filepath.Join(filepath.Dir(p.loadedPath), p.ImagesPath, image.Path)
	err := copyFile(originalImagePath, imagePath)
	if err != nil {
		return fmt.Errorf("error creating meeting for project %s: %w", p.Name, err)
	}

	p.Images = append(p.Images, image)
	return nil
}

func (p *Project) AddTag(tagName string) {
	p.Tags = append(p.Tags, &Tag{
		Created: time.Now(),
		Name:    tagName,
	})
}

func (p *Project) NewTask(name string) error {
	t := time.Now()
	task := &Task{
		Entry: Entry{
			Created: t,
			Path:    strconv.Itoa(len(p.Tasks) + 1),
		},
		Name:    name,
		History: make([]*TaskStatusChange, 0),
	}

	err := task.initiailzeMarkdown(filepath.Join(filepath.Dir(p.loadedPath), p.TasksPath))
	if err != nil {
		return fmt.Errorf("error creating task for project %s: %w", p.Name, err)
	}

	p.Tasks = append(p.Tasks, task)
	return nil
}

func (p *Project) NewTopic(name string) error {
	topic := &Topic{
		Entry: Entry{
			Created: time.Now(),
			Path:    removeNonAlphanumeric(name),
		},
		Name: name,
	}

	err := topic.initiailzeMarkdown(filepath.Join(filepath.Dir(p.loadedPath), p.TopicsPath))
	if err != nil {
		return fmt.Errorf("error creating topic for project %s: %w", p.Name, err)
	}

	p.Topics = append(p.Topics, topic)
	return nil
}

func (p *Project) Save() error {
	return SaveJSON(p.loadedPath, p)
}
