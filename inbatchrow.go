package rowmodels

import (
	"fmt"
	D "github.com/fbaube/dsmnd"
	FU "github.com/fbaube/fileutils"
	RU "github.com/fbaube/repoutils"
	// "github.com/fbaube/nurepo/db"
)

/*
var TableSummary_ContentityRow = D.TableSummary{
var TableDescriptor_ContentityRow = TableDescriptor{
func (cro *ContentityRow) PtrFields() []any { // barfs on []db.PtrFields
var ColumnSpecs_ContentityRow = []D.ColumnSpec{
type ContentityRow struct {
func (p *ContentityRow) String() string {
*/

// TableSummary_InbatchRow describes the table.
var TableSummary_InbatchRow = D.TableSummary{
	/*D.TABL*/ "TABLE", "INB", "inbatch", "Input batch of imported files"}

// TableDescriptor_InbatchRow TBS and no foreign keys.
var TableDescriptor_InbatchRow = RU.TableDescriptor{
	"inbatch",     // Name
	"inb",         // ShortName
	"idx_inbatch", // IDName
	//
	"FilCt, RelFP, AbsFP, T_Cre, Descr", // ColumnNames
	// No foreign keys
	ColumnSpecs_InbatchRow, // []D.ColumnSpec
}

// TODO: Still can't sort out the notation for ptr constraints ?!
func (inbro *InbatchRow) PtrFields() []any { // barfs on []db.PtrFields
	return []any{&inbro.Idx_Inbatch, &inbro.FilCt, &inbro.RelFP,
		&inbro.AbsFP, &inbro.T_Cre, &inbro.Descr}
}

// ColumnSpecs_InbatchRow specifies:
//   - file count
//   - two path fields (rel & abs)
//   - creation time
//   - description
//
// .
var ColumnSpecs_InbatchRow = []D.ColumnSpec{
	D.ColumnSpec{D.BDT_INTG, "filct",
		"Nr. of files", "Number of files"}, // D.INTEG
	D.DD_RelFP,
	D.DD_AbsFP,
	D.DD_T_Cre, // D.DATIM // THIS AND T_Edt SHOULD USE
	// DEFAULT https://www.sqlite.org/lang_createtable.html#dfltval
	// CURRENT_TIMESTAMP "YYYY-MM-DD HH:MM:SS"
	D.ColumnSpec{D.BDT_TEXT, "descr",
		"Batch descr.", "Inbatch description"}, // D.STRNG
}

// InbatchRow describes (in the DB)
// a single import batch at the CLI.
//   - NOTE: Maybe rename this to FileSet ?
//   - TODO: Maybe represent this with a dsmnd.NSPath: Batch.nr+Path
type InbatchRow struct {
	Idx_Inbatch int
	FilCt       int
	RelFP       string
	AbsFP       FU.AbsFilePath
	T_Cre       string
	// Batch "name" ? Or add a separate Name field ?
	Descr string
}

// TODO Write col desc's using Desmond !
// TODO Generate ColNames from ColumnSpecs_InbatchRow

// String implements Stringer. FIXME
func (p *InbatchRow) String() string {
	return fmt.Sprintf("inbatchrow FIXME")
	// p.PathProps.String(), p.PathAnalysis.String())
}
