package data_structures
// clarifying questions
// * need to handle "."? no

// approaches
// * go left or right
// * handle words or handle individual characters 
// * split string into components by "/"
// * key question is how handle ".." i think ideally want a stack where can go up and down
func simplifyPath(path string) string {
    // setup stack
    stack := []string{}

    // split folders
    folders := strings.Split(path, "/")
    // for loop through folders
    for _, f := range folders {
        // process different types
        switch f {
            case "":
                // do nothing
            case ".":
                // do nothing
            case "..":
                // pop stack
                if len(stack) > 0 {
                    stack = stack[:len(stack) - 1]
                }
            default:
                stack = append(stack, f)
        }
    }
    if len(stack) == 0 {
        return "/"
    }
    simplifiedPath := ""
    // build path - go backwards (it is a stack anyway)
    for i := len(stack) - 1; i >= 0; i-- {
        popped := stack[i]
        simplifiedPath = "/" + popped + simplifiedPath
    }
    return simplifiedPath
}