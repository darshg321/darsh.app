package main

// Mirrors the shape scripts/export-term.js writes.

type Link struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Project struct {
	Slug        string   `json:"slug"`
	Title       string   `json:"title"`
	Tagline     string   `json:"tagline"`
	When        string   `json:"when"`
	Year        int      `json:"year"`
	Kind        string   `json:"kind"`
	Status      string   `json:"status"`
	StatusLabel string   `json:"statusLabel"`
	Mask        string   `json:"mask"`
	Stack       []string `json:"stack"`
	With        []string `json:"with"`
	Events      []string `json:"events"`
	Award       string   `json:"award"`
	Repo        string   `json:"repo"`
	Links       []Link   `json:"links"`
	Body        string   `json:"body"`
}

type Stat struct {
	N     string `json:"n"`
	Label string `json:"label"`
}

type Event struct {
	Slug     string   `json:"slug"`
	Name     string   `json:"name"`
	Kind     string   `json:"kind"`
	Role     string   `json:"role"`
	When     string   `json:"when"`
	Place    string   `json:"place"`
	Theme    string   `json:"theme"`
	URL      string   `json:"url"`
	Summary  string   `json:"summary"`
	Stats    []Stat   `json:"stats"`
	Award    string   `json:"award"`
	Projects []string `json:"projects"`
}

type Role struct {
	Title  string   `json:"title"`
	Org    string   `json:"org"`
	When   string   `json:"when"`
	Points []string `json:"points"`
	Mask   string   `json:"mask"`
}

type Note struct {
	Title string `json:"title"`
	Note  string `json:"note"`
	When  string `json:"when"`
	URL   string `json:"url"`
}

type Content struct {
	Projects    []Project `json:"projects"`
	Events      []Event   `json:"events"`
	OtherEvents []struct {
		Name string `json:"name"`
		Note string `json:"note"`
		When string `json:"when"`
	} `json:"otherEvents"`
	Bench   []Note `json:"bench"`
	Early   []Note `json:"early"`
	Contact struct {
		Email string `json:"email"`
		Links []struct {
			Label  string `json:"label"`
			Handle string `json:"handle"`
			URL    string `json:"url"`
		} `json:"links"`
	} `json:"contact"`
	Roles []Role `json:"roles"`
	FRC   struct {
		Team    string   `json:"team"`
		Body    string   `json:"body"`
		Results []string `json:"results"`
	} `json:"frc"`
	Awards []struct {
		What string `json:"what"`
		Who  string `json:"who"`
		When string `json:"when"`
	} `json:"awards"`
	Offline []struct {
		What string `json:"what"`
		Note string `json:"note"`
	} `json:"offline"`
}

func (c *Content) project(slug string) *Project {
	for i := range c.Projects {
		if c.Projects[i].Slug == slug {
			return &c.Projects[i]
		}
	}
	return nil
}

func (c *Content) event(slug string) *Event {
	for i := range c.Events {
		if c.Events[i].Slug == slug {
			return &c.Events[i]
		}
	}
	return nil
}
