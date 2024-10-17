"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
(() => __awaiter(void 0, void 0, void 0, function* () {
    let iRootNode = {};
    let stSeparator = '/';
    const gstPaths = [];
    const loadJson = () => {
        return new Promise((resolve, reject) => {
            const xhr = new XMLHttpRequest();
            xhr.open('GET', 'dir.json', true);
            xhr.onreadystatechange = () => {
                if (4 !== xhr.readyState || 200 !== xhr.status) {
                    return;
                }
                if (200 === xhr.status) {
                    iRootNode = JSON.parse(xhr.responseText);
                    resolve('');
                }
            };
            xhr.send(null);
        });
    };
    const parseNode = (iNode, stParentDir) => {
        if (iNode.text !== '') {
            createDir(iNode, stParentDir);
        }
        if (stParentDir !== '') {
            stParentDir = stParentDir + stSeparator + iNode.text;
        }
        else {
            stParentDir = iNode.text;
        }
        if (iNode.children && iNode.children.length > 0) {
            for (const v of iNode.children) {
                parseNode(v, stParentDir);
            }
        }
    };
    const createDir = (iNode, stParentDir) => {
        var stDirPath = '';
        if (stParentDir !== '') {
            stDirPath = stParentDir + stSeparator + iNode.text;
        }
        else {
            stDirPath = iNode.text;
        }
        console.log(stDirPath);
    };
    yield loadJson();
    parseNode(iRootNode, '');
    console.log("test---");
}))();
