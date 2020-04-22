const glob = require("glob");
const markdownIt = require("markdown-it");
const meta = require("markdown-it-meta");
const fs = require("fs");
const _ = require("lodash");

const sidebar = (directory, array) => {
    return array.map(i => {
        const children = _.sortBy(
            glob
                .sync(`./${directory}/${i[1]}/*.md`)
                .map(path => {
                    const md = new markdownIt();
                    const file = fs.readFileSync(path, "utf8");
                    md.use(meta);
                    md.render(file);
                    const order = md.meta.order;
                    return { path, order };
                })
                .filter(f => f.order !== false),
            ["order", "path"]
        )
            .map(f => f.path)
            .filter(f => !f.match("README"));

        return {
            title: i[0],
            children
        };
    });
};

module.exports = {
    base: "/docs/",
    plugins: [
        ['@vuepress/search', {
            searchMaxSuggestions: 10
        }]
    ],
    locales: {
        "/": {
            lang: "简体中文",
            title: "IRITA 文档",
            description: "IRITA 文档",
        }
    },
    themeConfig: {
        repo: "bianjieai/irita",
        docsDir: "docs",
        editLinks: false,
        docsBranch: "master",
        locales: {
            "/": {
                sidebar: sidebar("", [
                    ["功能模块", "/features"],
                    ["命令行客户端", "/cli-client"],
                    ["API 服务", "/light-client"],
                ])
            }
        },
    }
};