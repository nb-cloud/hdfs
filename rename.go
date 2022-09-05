package hdfs

import (
	"os"

<<<<<<< HEAD
	hdfs "github.com/mizy/hdfs/internal/protocol/hadoop_hdfs"
=======
	hdfs "github.com/nb-cloud/hdfs/internal/protocol/hadoop_hdfs"
>>>>>>> 20c54cf (save didi hdfs proto)
	"google.golang.org/protobuf/proto"
)

// Rename renames (moves) a file.
func (c *Client) Rename(oldpath, newpath string) error {
	_, err := c.getFileInfo(newpath)
	err = interpretException(err)
	if err != nil && !os.IsNotExist(err) {
		return &os.PathError{"rename", newpath, err}
	}

	req := &hdfs.Rename2RequestProto{
		Src:           proto.String(oldpath),
		Dst:           proto.String(newpath),
		OverwriteDest: proto.Bool(true),
	}
	resp := &hdfs.Rename2ResponseProto{}

	err = c.namenode.Execute("rename2", req, resp)
	if err != nil {
		return &os.PathError{"rename", oldpath, interpretException(err)}
	}

	return nil
}
