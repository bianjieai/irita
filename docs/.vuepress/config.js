module.exports = {
  theme: "cosmos",
  title: "IRITA 技术文档",
  head: [
    [
      "link",
      {
        rel: "stylesheet",
        type: "text/css",
        href: "https://cloud.typography.com/6138116/7255612/css/fonts.css"
      }
    ],
  ],
  locales: {
    "/": {
      // lang: "en-US"
      lang: "cn"
    },
    // kr: {
    //   lang: "kr"
    // },
    cn: {
      lang: "cn"
    },
    // ru: {
    //   lang: "ru"
    // }
  },
  base: process.env.VUEPRESS_BASE || "/docs/",
  themeConfig: {
    repo: "irita/irita",
    docsRepo: "docs",
    docsDir: "/",
    editLinks: false,
    label: "irita",
    algolia: {
      id: "BH4D9OD16A",
      key: "ac317234e6a42074175369b2f42e9754",
      index: "irita-docs"
    },
    sidebar: [
      {
        // title: "Using the SDK",
        // children: [
        //   {
        //     title: "Modules",
        //     directory: true,
        //     path: "/modules"
        //   }
        // ]
      },
      {
        // title: "Resources",
        // children: [
        //   {
        //     title: "Tutorials",
        //     path: "https://tutorials.cosmos.network"
        //   },
        //   {
        //     title: "SDK API Reference",
        //     path: "https://godoc.org/github.com/cosmos/cosmos-sdk"
        //   },
        //   {
        //     title: "REST API Spec",
        //     path: "https://cosmos.network/rpc/"
        //   }
        // ]
      }
    ],
    footer: {
      logo: "/logo.jpg",
      services:[
        {
          img: "/irita_logo.png",
          url:"https://irita.bianjie.ai",
          text: "支持下一代分布式商业系统的企业级联盟链产品线"
        },
        {
          img: "/github_logo.png",
          url:"https://github.com/bianjieai/irita",
          text: "了解更多，请访问 GitHub 开源地址"
        }
      ],
      textLink: {
        text: "irita.bianjie.ai",
        url: "https://irita.bianjie.ai",
        target: "_blank"
      },
      // smallprint:
      //   "This website is maintained by Bianjie Inc.",
    }
  },
  plugins: [
    // [
    //   "@vuepress/google-analytics",
    //   {
    //     ga: "UA-51029217-12"
    //   }
    // ],
    // [
    //   "sitemap",
    //   {
    //     hostname: "https://docs.irita.io"
    //   }
    // ]
  ]
};
