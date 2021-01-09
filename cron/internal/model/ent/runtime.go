// Code generated by entc, DO NOT EDIT.

package ent

import (
	"graduate/cron/internal/model/ent/job"
	"graduate/cron/internal/model/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	jobFields := schema.Job{}.Fields()
	_ = jobFields
	// jobDescName is the schema descriptor for name field.
	jobDescName := jobFields[1].Descriptor()
	// job.DefaultName holds the default value on creation for the name field.
	job.DefaultName = jobDescName.Default.(string)
	// jobDescCreator is the schema descriptor for creator field.
	jobDescCreator := jobFields[2].Descriptor()
	// job.DefaultCreator holds the default value on creation for the creator field.
	job.DefaultCreator = jobDescCreator.Default.(string)
	// jobDescURL is the schema descriptor for url field.
	jobDescURL := jobFields[3].Descriptor()
	// job.DefaultURL holds the default value on creation for the url field.
	job.DefaultURL = jobDescURL.Default.(string)
	// jobDescSpec is the schema descriptor for spec field.
	jobDescSpec := jobFields[4].Descriptor()
	// job.DefaultSpec holds the default value on creation for the spec field.
	job.DefaultSpec = jobDescSpec.Default.(string)
	// jobDescComment is the schema descriptor for comment field.
	jobDescComment := jobFields[5].Descriptor()
	// job.DefaultComment holds the default value on creation for the comment field.
	job.DefaultComment = jobDescComment.Default.(string)
	// jobDescUpdater is the schema descriptor for updater field.
	jobDescUpdater := jobFields[6].Descriptor()
	// job.DefaultUpdater holds the default value on creation for the updater field.
	job.DefaultUpdater = jobDescUpdater.Default.(string)
	// jobDescMethod is the schema descriptor for method field.
	jobDescMethod := jobFields[7].Descriptor()
	// job.DefaultMethod holds the default value on creation for the method field.
	job.DefaultMethod = jobDescMethod.Default.(string)
	// jobDescBody is the schema descriptor for body field.
	jobDescBody := jobFields[8].Descriptor()
	// job.DefaultBody holds the default value on creation for the body field.
	job.DefaultBody = jobDescBody.Default.(string)
	// jobDescHeader is the schema descriptor for header field.
	jobDescHeader := jobFields[9].Descriptor()
	// job.DefaultHeader holds the default value on creation for the header field.
	job.DefaultHeader = jobDescHeader.Default.(string)
	// jobDescStoppable is the schema descriptor for stoppable field.
	jobDescStoppable := jobFields[10].Descriptor()
	// job.DefaultStoppable holds the default value on creation for the stoppable field.
	job.DefaultStoppable = jobDescStoppable.Default.(int)
	// jobDescStatus is the schema descriptor for status field.
	jobDescStatus := jobFields[11].Descriptor()
	// job.DefaultStatus holds the default value on creation for the status field.
	job.DefaultStatus = jobDescStatus.Default.(int)
	// jobDescID is the schema descriptor for id field.
	jobDescID := jobFields[0].Descriptor()
	// job.IDValidator is a validator for the "id" field. It is called by the builders before save.
	job.IDValidator = jobDescID.Validators[0].(func(int) error)
}