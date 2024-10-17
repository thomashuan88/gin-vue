(async() => {
    interface Node {
        text: string;
        children: Node[];
    }

    let iRootNode: Node = {} as Node
    let stSeparator: string = '/'
    const gstPaths: string[] = []

    const loadJson = () => {
        return new Promise((resolve, reject) => {
            const xhr = new XMLHttpRequest()
            xhr.open('GET', 'dir.json', true)
            xhr.onreadystatechange = () => {
                if (4 !== xhr.readyState || 200 !== xhr.status) {
                    return
                }
                if (200 === xhr.status) {
                    iRootNode = JSON.parse(xhr.responseText) as Node
                    resolve('')
                }
            }
            xhr.send(null)
        })
    }

    const parseNode = (iNode: Node, stParentDir: string) => {
        if (iNode.text !== '') {
            createDir(iNode, stParentDir)
        }
        if (stParentDir !== '') {
            stParentDir = stParentDir + stSeparator + iNode.text
        } else {
            stParentDir = iNode.text
        }

        if (iNode.children && iNode.children.length > 0) {
            for (const v of iNode.children) {
                parseNode(v, stParentDir)
            }
        }
    }


    const createDir = (iNode: Node, stParentDir: string) => {
        var stDirPath = ''
        if (stParentDir !== '') {
            stDirPath = stParentDir + stSeparator + iNode.text
        } else {
            stDirPath = iNode.text
        }
        console.log(stDirPath)

    }

    await loadJson()

    parseNode(iRootNode, '')

    console.log("test---")
})()