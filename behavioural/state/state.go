package main

import "fmt"

type IState interface {
	Action(document *Document)
	GetName() string
}

type Document struct {
	state      IState
	isApproved bool
}

func NewDocument() *Document {
	return &Document{
		state:      NewDraftState(),
		isApproved: false,
	}
}

func (d *Document) GetState() IState {
	return d.state
}

func (d *Document) SetState(state IState) {
	d.state = state
}

func (d *Document) SetApproval(bool bool) {
	d.isApproved = bool
}

func (d *Document) Publish() {
	d.state.Action(d)
}

type DraftState struct {
	name string
}

func NewDraftState() *DraftState {
	return &DraftState{name: "Draft"}
}

func (ds *DraftState) GetName() string {
	return ds.name
}

func (ds *DraftState) Action(document *Document) {
	document.SetState(NewReviewState())
}

type ReviewState struct {
	name string
}

func NewReviewState() *ReviewState {
	return &ReviewState{name: "Review"}
}

func (rs *ReviewState) GetName() string {
	return rs.name
}

func (rs *ReviewState) Action(document *Document) {
	if document.isApproved {
		document.SetState(NewPublishedState())
		return
	}
	document.SetState(NewDraftState())
}

type PublishedState struct {
	name string
}

func NewPublishedState() *PublishedState {
	return &PublishedState{name: "Published"}
}

func (ps *PublishedState) GetName() string {
	return ps.name
}

func (ps *PublishedState) Action(document *Document) {
	fmt.Println("Document already published, no change anymore !")
}

func main() {
	doc := NewDocument()
	fmt.Println(doc.GetState().GetName())

	doc.Publish()
	fmt.Println(doc.GetState().GetName())

	doc.Publish()
	fmt.Println(doc.GetState().GetName())

	doc.SetApproval(true)
	doc.Publish()
	doc.Publish()
	fmt.Println(doc.GetState().GetName())

	doc.Publish()
	fmt.Println(doc.GetState().GetName())
}
