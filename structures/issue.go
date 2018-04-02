package structures

import (
	"github.com/trivago/tgo/tcontainer"
	"time"
	"reflect"
	"strings"
	"fmt"
	"encoding/json"
)

type Issue struct {
	Expand    string       `json:"expand,omitempty" structs:"expand,omitempty"`
	ID        string       `json:"id,omitempty" structs:"id,omitempty"`
	Self      string       `json:"self,omitempty" structs:"self,omitempty"`
	Key       string       `json:"key,omitempty" structs:"key,omitempty"`
	Fields    *IssueFields `json:"fields,omitempty" structs:"fields,omitempty"`
	Changelog *Changelog   `json:"changelog,omitempty" structs:"changelog,omitempty"`
}

type ChangelogItems struct {
	Field      string      `json:"field" structs:"field"`
	FieldType  string      `json:"fieldtype" structs:"fieldtype"`
	From       interface{} `json:"from" structs:"from"`
	FromString string      `json:"fromString" structs:"fromString"`
	To         interface{} `json:"to" structs:"to"`
	ToString   string      `json:"toString" structs:"toString"`
}

type ChangelogHistory struct {
	Id      string           `json:"id" structs:"id"`
	Author  User             `json:"author" structs:"author"`
	Created string           `json:"created" structs:"created"`
	Items   []ChangelogItems `json:"items" structs:"items"`
}

type Changelog struct {
	Histories []ChangelogHistory `json:"histories,omitempty"`
}

type Attachment struct {
	Self      string `json:"self,omitempty" structs:"self,omitempty"`
	ID        string `json:"id,omitempty" structs:"id,omitempty"`
	Filename  string `json:"filename,omitempty" structs:"filename,omitempty"`
	Author    *User  `json:"author,omitempty" structs:"author,omitempty"`
	Created   string `json:"created,omitempty" structs:"created,omitempty"`
	Size      int    `json:"size,omitempty" structs:"size,omitempty"`
	MimeType  string `json:"mimeType,omitempty" structs:"mimeType,omitempty"`
	Content   string `json:"content,omitempty" structs:"content,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty" structs:"thumbnail,omitempty"`
}

type Epic struct {
	ID      int    `json:"id" structs:"id"`
	Key     string `json:"key" structs:"key"`
	Self    string `json:"self" structs:"self"`
	Name    string `json:"name" structs:"name"`
	Summary string `json:"summary" structs:"summary"`
	Done    bool   `json:"done" structs:"done"`
}

type IssueFields struct {
	// TODO Missing fields
	//      * "aggregatetimespent": null,
	//      * "workratio": -1,
	//      * "lastViewed": null,
	//      * "aggregatetimeoriginalestimate": null,
	//      * "aggregatetimeestimate": null,
	//      * "environment": null,
	Expand               string        `json:"expand,omitempty" structs:"expand,omitempty"`
	Type                 IssueType     `json:"issuetype,omitempty" structs:"issuetype,omitempty"`
	Project              Project       `json:"project,omitempty" structs:"project,omitempty"`
	Resolution           *Resolution   `json:"resolution,omitempty" structs:"resolution,omitempty"`
	Priority             *Priority     `json:"priority,omitempty" structs:"priority,omitempty"`
	Resolutiondate       Time          `json:"resolutiondate,omitempty" structs:"resolutiondate,omitempty"`
	Created              Time          `json:"created,omitempty" structs:"created,omitempty"`
	Duedate              Date          `json:"duedate,omitempty" structs:"duedate,omitempty"`
	Watches              *Watches      `json:"watches,omitempty" structs:"watches,omitempty"`
	Assignee             *User         `json:"assignee,omitempty" structs:"assignee,omitempty"`
	Updated              Time          `json:"updated,omitempty" structs:"updated,omitempty"`
	Description          string        `json:"description,omitempty" structs:"description,omitempty"`
	Summary              string        `json:"summary,omitempty" structs:"summary,omitempty"`
	Creator              *User         `json:"Creator,omitempty" structs:"Creator,omitempty"`
	Reporter             *User         `json:"reporter,omitempty" structs:"reporter,omitempty"`
	Components           []*Component  `json:"components,omitempty" structs:"components,omitempty"`
	Status               *Status       `json:"status,omitempty" structs:"status,omitempty"`
	Progress             *Progress     `json:"progress,omitempty" structs:"progress,omitempty"`
	AggregateProgress    *Progress     `json:"aggregateprogress,omitempty" structs:"aggregateprogress,omitempty"`
	TimeTracking         *TimeTracking `json:"timetracking,omitempty" structs:"timetracking,omitempty"`
	TimeSpent            int           `json:"timespent,omitempty" structs:"timespent,omitempty"`
	TimeEstimate         int           `json:"timeestimate,omitempty" structs:"timeestimate,omitempty"`
	TimeOriginalEstimate int           `json:"timeoriginalestimate,omitempty" structs:"timeoriginalestimate,omitempty"`
	Worklog              *Worklog      `json:"worklog,omitempty" structs:"worklog,omitempty"`
	IssueLinks           []*IssueLink  `json:"issuelinks,omitempty" structs:"issuelinks,omitempty"`
	Comments             *Comments     `json:"comment,omitempty" structs:"comment,omitempty"`
	FixVersions          []*FixVersion `json:"fixVersions,omitempty" structs:"fixVersions,omitempty"`
	Labels               []string      `json:"labels,omitempty" structs:"labels,omitempty"`
	Subtasks             []*Subtasks   `json:"subtasks,omitempty" structs:"subtasks,omitempty"`
	Attachments          []*Attachment `json:"attachment,omitempty" structs:"attachment,omitempty"`
	Epic                 *Epic         `json:"epic,omitempty" structs:"epic,omitempty"`
	Parent               *Parent       `json:"parent,omitempty" structs:"parent,omitempty"`
	Unknowns             tcontainer.MarshalMap
}


type IssueType struct {
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	ID          string `json:"id,omitempty" structs:"id,omitempty"`
	Description string `json:"description,omitempty" structs:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty" structs:"iconUrl,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	Subtask     bool   `json:"subtask,omitempty" structs:"subtask,omitempty"`
	AvatarID    int    `json:"avatarId,omitempty" structs:"avatarId,omitempty"`
}

// Resolution represents a resolution of a JIRA issue.
// Typical types are "Fixed", "Suspended", "Won't Fix", ...
type Resolution struct {
	Self        string `json:"self" structs:"self"`
	ID          string `json:"id" structs:"id"`
	Description string `json:"description" structs:"description"`
	Name        string `json:"name" structs:"name"`
}

// Priority represents a priority of a JIRA issue.
// Typical types are "Normal", "Moderate", "Urgent", ...
type Priority struct {
	Self    string `json:"self,omitempty" structs:"self,omitempty"`
	IconURL string `json:"iconUrl,omitempty" structs:"iconUrl,omitempty"`
	Name    string `json:"name,omitempty" structs:"name,omitempty"`
	ID      string `json:"id,omitempty" structs:"id,omitempty"`
}

// Watches represents a type of how many and which user are "observing" a JIRA issue to track the status / updates.
type Watches struct {
	Self       string     `json:"self,omitempty" structs:"self,omitempty"`
	WatchCount int        `json:"watchCount,omitempty" structs:"watchCount,omitempty"`
	IsWatching bool       `json:"isWatching,omitempty" structs:"isWatching,omitempty"`
	Watchers   []*Watcher `json:"watchers,omitempty" structs:"watchers,omitempty"`
}

// Watcher represents a simplified user that "observes" the issue
type Watcher struct {
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active      bool   `json:"active,omitempty" structs:"active,omitempty"`
}

// AvatarUrls represents different dimensions of avatars / images
type AvatarUrls struct {
	Four8X48  string `json:"48x48,omitempty" structs:"48x48,omitempty"`
	Two4X24   string `json:"24x24,omitempty" structs:"24x24,omitempty"`
	One6X16   string `json:"16x16,omitempty" structs:"16x16,omitempty"`
	Three2X32 string `json:"32x32,omitempty" structs:"32x32,omitempty"`
}

// Component represents a "component" of a JIRA issue.
// Components can be user defined in every JIRA instance.
type Component struct {
	Self string `json:"self,omitempty" structs:"self,omitempty"`
	ID   string `json:"id,omitempty" structs:"id,omitempty"`
	Name string `json:"name,omitempty" structs:"name,omitempty"`
}

// Status represents the current status of a JIRA issue.
// Typical status are "Open", "In Progress", "Closed", ...
// Status can be user defined in every JIRA instance.
type Status struct {
	Self           string         `json:"self" structs:"self"`
	Description    string         `json:"description" structs:"description"`
	IconURL        string         `json:"iconUrl" structs:"iconUrl"`
	Name           string         `json:"name" structs:"name"`
	ID             string         `json:"id" structs:"id"`
	StatusCategory StatusCategory `json:"statusCategory" structs:"statusCategory"`
}

// StatusCategory represents the category a status belongs to.
// Those categories can be user defined in every JIRA instance.
type StatusCategory struct {
	Self      string `json:"self" structs:"self"`
	ID        int    `json:"id" structs:"id"`
	Name      string `json:"name" structs:"name"`
	Key       string `json:"key" structs:"key"`
	ColorName string `json:"colorName" structs:"colorName"`
}

// Progress represents the progress of a JIRA issue.
type Progress struct {
	Progress int `json:"progress" structs:"progress"`
	Total    int `json:"total" structs:"total"`
}

// Parent represents the parent of a JIRA issue, to be used with subtask issue types.
type Parent struct {
	ID  string `json:"id,omitempty" structs:"id"`
	Key string `json:"key,omitempty" structs:"key"`
}

// Time represents the Time definition of JIRA as a time.Time of go
type Time time.Time

// Date represents the Date definition of JIRA as a time.Time of go
type Date time.Time

// Wrapper struct for search result
type transitionResult struct {
	Transitions []Transition `json:"transitions" structs:"transitions"`
}

// Transition represents an issue transition in JIRA
type Transition struct {
	ID     string                     `json:"id" structs:"id"`
	Name   string                     `json:"name" structs:"name"`
	To     Status                     `json:"to" structs:"status"`
	Fields map[string]TransitionField `json:"fields" structs:"fields"`
}

// TransitionField represents the value of one Transition
type TransitionField struct {
	Required bool `json:"required" structs:"required"`
}

// CreateTransitionPayload is used for creating new issue transitions
type CreateTransitionPayload struct {
	Transition TransitionPayload       `json:"transition" structs:"transition"`
	Fields     TransitionPayloadFields `json:"fields" structs:"fields"`
}

// TransitionPayload represents the request payload of Transition calls like DoTransition
type TransitionPayload struct {
	ID string `json:"id" structs:"id"`
}

// TransitionPayloadFields represents the fields that can be set when executing a transition
type TransitionPayloadFields struct {
	Resolution *Resolution `json:"resolution,omitempty" structs:"resolution,omitempty"`
}

// Option represents an option value in a SelectList or MultiSelect
// custom issue field
type Option struct {
	Value string `json:"value" structs:"value"`
}

type Worklog struct {
	StartAt    int             `json:"startAt" structs:"startAt"`
	MaxResults int             `json:"maxResults" structs:"maxResults"`
	Total      int             `json:"total" structs:"total"`
	Worklogs   []WorklogRecord `json:"worklogs" structs:"worklogs"`
}

// WorklogRecord represents one entry of a Worklog
type WorklogRecord struct {
	Self             string `json:"self,omitempty" structs:"self,omitempty"`
	Author           *User  `json:"author,omitempty" structs:"author,omitempty"`
	UpdateAuthor     *User  `json:"updateAuthor,omitempty" structs:"updateAuthor,omitempty"`
	Comment          string `json:"comment,omitempty" structs:"comment,omitempty"`
	Created          *Time  `json:"created,omitempty" structs:"created,omitempty"`
	Updated          *Time  `json:"updated,omitempty" structs:"updated,omitempty"`
	Started          *Time  `json:"started,omitempty" structs:"started,omitempty"`
	TimeSpent        string `json:"timeSpent,omitempty" structs:"timeSpent,omitempty"`
	TimeSpentSeconds int    `json:"timeSpentSeconds,omitempty" structs:"timeSpentSeconds,omitempty"`
	ID               string `json:"id,omitempty" structs:"id,omitempty"`
	IssueID          string `json:"issueId,omitempty" structs:"issueId,omitempty"`
}

// TimeTracking represents the timetracking fields of a JIRA issue.
type TimeTracking struct {
	OriginalEstimate         string `json:"originalEstimate,omitempty" structs:"originalEstimate,omitempty"`
	RemainingEstimate        string `json:"remainingEstimate,omitempty" structs:"remainingEstimate,omitempty"`
	TimeSpent                string `json:"timeSpent,omitempty" structs:"timeSpent,omitempty"`
	OriginalEstimateSeconds  int    `json:"originalEstimateSeconds,omitempty" structs:"originalEstimateSeconds,omitempty"`
	RemainingEstimateSeconds int    `json:"remainingEstimateSeconds,omitempty" structs:"remainingEstimateSeconds,omitempty"`
	TimeSpentSeconds         int    `json:"timeSpentSeconds,omitempty" structs:"timeSpentSeconds,omitempty"`
}

// Subtasks represents all issues of a parent issue.
type Subtasks struct {
	ID     string      `json:"id" structs:"id"`
	Key    string      `json:"key" structs:"key"`
	Self   string      `json:"self" structs:"self"`
	Fields IssueFields `json:"fields" structs:"fields"`
}

// IssueLink represents a link between two issues in JIRA.
type IssueLink struct {
	ID           string        `json:"id,omitempty" structs:"id,omitempty"`
	Self         string        `json:"self,omitempty" structs:"self,omitempty"`
	Type         IssueLinkType `json:"type" structs:"type"`
	OutwardIssue *Issue        `json:"outwardIssue" structs:"outwardIssue"`
	InwardIssue  *Issue        `json:"inwardIssue" structs:"inwardIssue"`
	Comment      *Comment      `json:"comment,omitempty" structs:"comment,omitempty"`
}

// IssueLinkType represents a type of a link between to issues in JIRA.
// Typical issue link types are "Related to", "Duplicate", "Is blocked by", etc.
type IssueLinkType struct {
	ID      string `json:"id,omitempty" structs:"id,omitempty"`
	Self    string `json:"self,omitempty" structs:"self,omitempty"`
	Name    string `json:"name" structs:"name"`
	Inward  string `json:"inward" structs:"inward"`
	Outward string `json:"outward" structs:"outward"`
}

// Comments represents a list of Comment.
type Comments struct {
	Comments []*Comment `json:"comments,omitempty" structs:"comments,omitempty"`
}

// Comment represents a comment by a person to an issue in JIRA.
type Comment struct {
	ID           string            `json:"id,omitempty" structs:"id,omitempty"`
	Self         string            `json:"self,omitempty" structs:"self,omitempty"`
	Name         string            `json:"name,omitempty" structs:"name,omitempty"`
	Author       User              `json:"author,omitempty" structs:"author,omitempty"`
	Body         string            `json:"body,omitempty" structs:"body,omitempty"`
	UpdateAuthor User              `json:"updateAuthor,omitempty" structs:"updateAuthor,omitempty"`
	Updated      string            `json:"updated,omitempty" structs:"updated,omitempty"`
	Created      string            `json:"created,omitempty" structs:"created,omitempty"`
	Visibility   CommentVisibility `json:"visibility,omitempty" structs:"visibility,omitempty"`
}

// FixVersion represents a software release in which an issue is fixed.
type FixVersion struct {
	Archived        *bool  `json:"archived,omitempty" structs:"archived,omitempty"`
	ID              string `json:"id,omitempty" structs:"id,omitempty"`
	Name            string `json:"name,omitempty" structs:"name,omitempty"`
	ProjectID       int    `json:"projectId,omitempty" structs:"projectId,omitempty"`
	ReleaseDate     string `json:"releaseDate,omitempty" structs:"releaseDate,omitempty"`
	Released        *bool  `json:"released,omitempty" structs:"released,omitempty"`
	Self            string `json:"self,omitempty" structs:"self,omitempty"`
	UserReleaseDate string `json:"userReleaseDate,omitempty" structs:"userReleaseDate,omitempty"`
}

// CommentVisibility represents he visibility of a comment.
// E.g. Type could be "role" and Value "Administrators"
type CommentVisibility struct {
	Type  string `json:"type,omitempty" structs:"type,omitempty"`
	Value string `json:"value,omitempty" structs:"value,omitempty"`
}

// SearchOptions specifies the optional parameters to various List methods that
// support pagination.
// Pagination is used for the JIRA REST APIs to conserve server resources and limit
// response size for resources that return potentially large collection of items.
// A request to a pages API will result in a values array wrapped in a JSON object with some paging metadata
// Default Pagination options
type SearchOptions struct {
	// StartAt: The starting index of the returned projects. Base index: 0.
	StartAt int `url:"startAt,omitempty"`
	// MaxResults: The maximum number of projects to return per page. Default: 50.
	MaxResults int `url:"maxResults,omitempty"`
	// Expand: Expand specific sections in the returned issues
	Expand string `url:"expand,omitempty"`
	Fields []string
	// ValidateQuery: The validateQuery param offers control over whether to validate and how strictly to treat the validation. Default: strict.
	ValidateQuery string `url:"validateQuery,omitempty"`
}

// searchResult is only a small wrapper around the Search (with JQL) method
// to be able to parse the results
type searchResult struct {
	Issues     []Issue `json:"issues" structs:"issues"`
	StartAt    int     `json:"startAt" structs:"startAt"`
	MaxResults int     `json:"maxResults" structs:"maxResults"`
	Total      int     `json:"total" structs:"total"`
}

// GetQueryOptions specifies the optional parameters for the Get Issue methods
type GetQueryOptions struct {
	// Fields is the list of fields to return for the issue. By default, all fields are returned.
	Fields string `url:"fields,omitempty"`
	Expand string `url:"expand,omitempty"`
	// Properties is the list of properties to return for the issue. By default no properties are returned.
	Properties string `url:"properties,omitempty"`
	// FieldsByKeys if true then fields in issues will be referenced by keys instead of ids
	FieldsByKeys  bool   `url:"fieldsByKeys,omitempty"`
	UpdateHistory bool   `url:"updateHistory,omitempty"`
	ProjectKeys   string `url:"projectKeys,omitempty"`
}

// CustomFields represents custom fields of JIRA
// This can heavily differ between JIRA instances
type CustomFields map[string]string
////////////////////////////////////////////////////////////////////
func (i *IssueFields) UnmarshalJSON(data []byte) error {

	// Do the normal unmarshalling first
	// Details for this way: http://choly.ca/post/go-json-marshalling/
	type Alias IssueFields
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	totalMap := tcontainer.NewMarshalMap()
	err := json.Unmarshal(data, &totalMap)
	if err != nil {
		return err
	}

	t := reflect.TypeOf(*i)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagDetail := field.Tag.Get("json")
		if tagDetail == "" {
			// ignore if there are no tags
			continue
		}
		options := strings.Split(tagDetail, ",")

		if len(options) == 0 {
			return fmt.Errorf("No tags options found for %s", field.Name)
		}
		// the first one is the json tag
		key := options[0]
		if _, okay := totalMap.Value(key); okay {
			delete(totalMap, key)
		}

	}
	i = (*IssueFields)(aux.Alias)
	// all the tags found in the struct were removed. Whatever is left are unknowns to struct
	i.Unknowns = totalMap
	return nil

}