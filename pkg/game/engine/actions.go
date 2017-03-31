package engine

import (
	"strings"

	//fusefs "bazil.org/fuse/fs"

	sh "github.com/ChrisRx/dungeonfs/pkg/commands"
	"github.com/ChrisRx/dungeonfs/pkg/game/assets"
	"github.com/ChrisRx/dungeonfs/pkg/game/fs"
)

//func createAction(name string, node fs.Node) fusefs.Node {
func createAction(name string, node fs.Node) fs.Node {
	if node.Name() == "door" {
		if name == "wall" {
			node.MetaData().Set("Description", "You found a small switch on the wall and it opened up a path to the east.")
			newDir := node.New(fs.DirNode, "east")
			newDir.MetaData().Set("Description", "This place sucks")
			newDir.New(fs.FileNode, "<[TROLL]>").MetaData().Set("Content", []byte(assets.Troll))
			for _, f := range node.Children().Files() {
				return f
			}
		}
	}
	return node.New(fs.FileNode, name)
}

func parseArgs(s string) string {
	args := strings.Fields(s)
	if len(args) > 0 {
		return args[0]
	}
	return s
}

func lookupAction(name string, node fs.Node) fs.Node {
	switch parseArgs(name) {
	case "look":
		desc := node.MetaData().GetString("Description")
		f := node.New(fs.FileNode, "look")
		f.MetaData().Set("Content", sh.Script(sh.Echo(desc)))
		return f
	case "sword":
		for _, f := range node.Children().Files() {
			commands := []string{
				sh.Echo("you swing your sword mightily at the %s ...", f.Name()),
				sh.Command("sleep 1"),
				sh.Echo("It appeared to have no effect."),
			}
			f := node.New(fs.FileNode, "sword")
			f.MetaData().Set("Content", sh.Script(commands...))
			return f
		}
	}
	return nil
}