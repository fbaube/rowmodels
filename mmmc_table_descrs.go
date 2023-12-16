package rowmodels

import RU "github.com/fbaube/repoutils"

// MmmcTableDescriptors configures the three key tables.
var MmmcTableDescriptors = []RU.TableDescriptor{
	TableDescriptor_InbatchRow,
	TableDescriptor_ContentityRow,
	TableDescriptor_TopicrefRow,
}
