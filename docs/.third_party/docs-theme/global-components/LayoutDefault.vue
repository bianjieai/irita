<template lang="pug">
  div(style="width: 100%")
    .search__container
      .search(@click="$emit('search', true)")
        .search__icon
          icon-search
        .search__text Search
    .container
      slot
      tm-content-cards(v-if="$frontmatter.cards")
</template>

<style lang="stylus" scoped>
.search
  display flex
  align-items center
  color rgba(22, 25, 49, 0.65)
  padding-top 0.5rem
  width calc(var(--aside-width) - 6rem)
  cursor pointer
  position absolute
  top 1rem
  right 4rem
  justify-content flex-end

  &__container
    visibility hidden
    display flex
    justify-content flex-end
    margin-top 1rem
    margin-bottom 1rem

  &__icon
    width 1.5rem
    height 1.5rem
    fill #aaa
    margin-right 0.5rem

.footer__links
  padding-top 5rem
  padding-bottom 1rem
  border-top 1px solid rgba(176, 180, 207, 0.2)
  margin-top 5rem

.links
  display flex
  justify-content space-between
  margin-top 4rem

  a
    box-shadow none
    color var(--accent-color)

.container
  position relative
  min-height 100vh
  width 100%

.content
  padding-right var(--sidebar-width)
  width 100%
  position relative

  &.noAside
    padding-right 0

  &__container
    width 100%
    padding-left 4rem
    padding-right 2rem

    &.noAside
      max-width initial

/deep/
  .codeblock
    margin-top 1rem
    margin-bottom 1rem

  .custom-block
    &.danger, &.warning, &.tip
      padding 1rem 1.5rem 1rem 3.5rem
      border-radius 0.5rem
      position relative

      & :first-child
        margin-top 0

      & :last-child
        margin-bottom 0

      &:before
        content ''
        height 24px
        width 24px
        position absolute
        display block
        top 1rem
        left 1rem
        background-repeat no-repeat

    &.danger
      background #FFF6F9

      &:before
        background-image url("./images/icon-danger.svg")

    &.warning
      &:before
        background-image url("./images/icon-warning.svg")

    &.tip
      &:before
        background-image url("./images/icon-tip.svg")

  h2, h3, h4, h5, h6
    &:hover
      a.header-anchor
        opacity 1

  a.header-anchor
    opacity 0
    position absolute
    font-weight 400
    left -1.5em
    width 1.5em
    text-align center
    box-sizing border-box
    color rgba(0, 0, 0, 0.5)
    transition all 0.25s

    &:after
      transition all 0.25s
      border-radius 0.25rem
      content attr(data-header-anchor-text)
      width 4rem
      color white
      position absolute
      top -2.4em
      padding 7px 12px
      white-space nowrap
      left 50%
      transform translateX(-50%)
      font-size 0.8125rem
      line-height 1
      opacity 0
      box-shadow 0px 16px 32px rgba(22, 25, 49, 0.08), 0px 8px 12px rgba(22, 25, 49, 0.06), 0px 1px 0px rgba(22, 25, 49, 0.05)
      background #161931

    &:before
      transition all 0.25s
      content ''
      background-image url("data:image/svg+xml,  <svg xmlns='http://www.w3.org/2000/svg' width='100%' height='100%' viewBox='0 0 24 24'><path fill='rgb(22, 25, 49)' d='M12 21l-12-18h24z'/></svg>")
      position absolute
      width 8px
      height 8px
      top -0.7em
      left 50%
      font-size 0.5rem
      transform translateX(-50%)
      opacity 0

    &:hover:before
      opacity 1

    &:hover:after
      opacity 1

  h1[id*='requisite'], h2[id*='requisite'], h3[id*='requisite'], h4[id*='requisite'], h5[id*='requisite'], h6[id*='requisite']
    display none
    align-items center
    cursor pointer

    &:before
      content ''
      width 24px
      height 24px
      display block
      margin-right 0.75rem
      background url('./images/icon-chevron.svg')
      transition all 0.25s

  h1[id*='requisite'].prereqTitleShow, h2[id*='requisite'].prereqTitleShow, h3[id*='requisite'].prereqTitleShow, h4[id*='requisite'].prereqTitleShow, h5[id*='requisite'].prereqTitleShow, h6[id*='requisite'].prereqTitleShow
    &:before
      transform rotate(90deg)

  h1[id*='requisite'] + ul, h2[id*='requisite'] + ul, h3[id*='requisite'] + ul, h4[id*='requisite'] + ul, h5[id*='requisite'] + ul, h6[id*='requisite'] + ul
    padding-left initial
    display none

  li[prereq]
    padding-left initial
    display none

    &:before
      display none

  li[prereq].prereqLinkShow
    display block

  li[prereq] a[href]
    box-shadow 0px 2px 4px rgba(22, 25, 49, 0.05), 0px 0px 1px rgba(22, 25, 49, 0.2), 0px 0.5px 0px rgba(22, 25, 49, 0.05)
    padding 1rem
    border-radius 0.5rem
    color #161931
    font-size 0.875rem
    font-weight 500
    line-height 20px
    margin 1rem 0
    display block
    letter-spacing 0.01em

    &:hover
      color inherit

  [synopsis]
    padding 1.5rem 2rem
    background-color rgba(176, 180, 207, 0.09)
    border-radius 0.5rem
    margin-top 3rem
    margin-bottom 3rem
    letter-spacing 0.01em
    color rgba(22, 25, 49, 0.9)
    font-size 0.875rem
    line-height 20px

    &:before
      content 'Synopsis'
      display block
      color rgba(22, 25, 49, 0.65)
      text-transform uppercase
      font-size 0.75rem
      margin-bottom 0.5rem
      letter-spacing 0.2em

  [synopsis]
    & a
      box-shadow none

    & a:hover
      box-shadow none

    & a code
      box-shadow 0 1px 0 0 rgba(80, 100, 251, 0.3), 0 0 0 3px #f8f8fb

    & a:hover code
      box-shadow 0 1px 0 0 rgba(80, 100, 251, 1), 0 0 0 3px #f8f8fb

    & a:active code
      color rgba(80, 100, 251, 0.6)
      box-shadow 0 1px 0 0 rgba(80, 100, 251, 0.3), 0 0 0 3px #f8f8fb

  a[target='_blank']:after
    content '↗'
    position absolute
    top 50%
    transform translateY(-50%)
    padding-left 3px
    font-size 0.75rem

  .icon.outbound
    display none

  table
    width 100%
    line-height 24px
    margin-top 2rem
    margin-bottom 2rem
    box-shadow 0 0 0 1px rgba(140, 145, 177, 0.32)
    border-radius 0.5rem
    border-collapse collapse

  td
    word-break break-word

  th
    text-align left
    font-weight 600
    font-size 0.875rem

  td, th
    padding 0.75rem

  tr
    box-shadow 0 1px 0 0 rgba(140, 145, 177, 0.32)

  tr:only-child
    box-shadow none

  thead tr:only-child
    box-shadow 0 1px 0 0 rgba(140, 145, 177, 0.32)

  tr + tr:last-child
    box-shadow none

  tr:last-child td
    border-bottom none

  .code-block__container
    margin-top 2rem
    margin-bottom 2rem

  .content__default
    width 100%

  h1, h2, h3, h4
    font-weight 600

  h1 code, h2 code, h3 code
    font-weight normal

  .content__container
    img
      max-width 100%

  .term
    text-decoration underline

  img
    width 100%
    height auto
    display block

  .tooltip
    h1
      font-size 0.875rem
      font-weight 500
      margin-bottom 0

    p
      margin-top 0
      margin-bottom 0
      line-height 1.5

  strong
    font-weight 600

  em
    font-style italic

  h1
    font-size 2.5rem
    font-weight 600
    margin-bottom 3rem
    line-height 3.25rem
    letter-spacing -0.03em

  h2
    font-size 1.5rem
    font-weight 600
    margin-top 3rem
    margin-bottom 1.5rem
    line-height 2rem
    letter-spacing -0.01em

  h3
    font-size 1.25rem
    font-weight 600
    margin-top 2rem
    margin-bottom 1rem
    letter-spacing -0.01em

  p
    margin-top 1rem
    margin-bottom 1rem
    line-height 1.5rem

  ul, ol
    line-height 1.5
    margin-top 1rem
    padding-left 0.75rem
    margin-bottom 1.5rem

  li
    padding-left 2rem
    list-style none
    margin-bottom 1rem
    position relative

    &:before
      content ''
      width 1rem
      height 1rem
      background url('./images/bullet-list.svg') no-repeat top left
      position absolute
      top 0.35rem
      left 0

  code
    background-color rgba(176, 180, 207, 0.2)
    border 1px solid rgba(176, 180, 207, 0.09)
    border-radius 0.25rem
    padding-left 0.25rem
    padding-right 0.25rem
    font-size 0.8125rem
    font-family 'Menlo', 'Monaco', 'Fira Code', monospace
    color #46509F
    margin-top 3rem

  h1, h2, h3, h4, h5, h6
    code
      font-size inherit

  p, ul, ol
    a
      color var(--accent-color)
      box-shadow inset 0 -0.5px 0 var(--accent-color)
      outline none
      transition box-shadow 0.25s
      line-height 12px
      position relative

    a[target='_blank']
      margin-right 1rem

    a:focus
      box-shadow 0 0 0 3px rgba(102, 161, 255, 0.7)
      border-radius 0.25rem

    a:focus:hover
      border-radius 0

    a:hover
      box-shadow inset 0 -1px 0 inherit

    a:active
      color var(--accent-color)
      border-radius 0
      box-shadow inset 0 -0.5px 0 var(--accent-color)
      opacity 0.65

    a code
      border-bottom none
      box-shadow 0 1px 0 0 rgba(80, 100, 251, 0.3), 0 0 0 3px white
      color rgba(80, 100, 251, 1)

    a:hover code
      box-shadow 0 1px 0 0 rgba(80, 100, 251, 1), 0 0 0 3px white

    a:active code
      color rgba(80, 100, 251, 0.6)
      box-shadow 0 1px 0 0 rgba(80, 100, 251, 0.3), 0 0 0 3px white
      background-color rgba(176, 180, 207, 0.1)
      border 1px solid rgba(176, 180, 207, 0.045)
      border-bottom none
      border-radius 0.25rem

    a:focus code
      box-shadow none

@media screen and (max-width: 1136px)
  >>> h2, >>> h3, >>> h4, >>> h5, >>> h6
    padding-right 1.5rem

  >>> a.header-anchor
    left initial
    right 0
    opacity 1

    &:after
      transform none
      left initial
      right -5px

  >>> h1 a.header-anchor
    display none

@media screen and (max-width: 1024px)
  .content
    padding-right 0

    &__container
      padding-left 2rem

@media screen and (max-width: 1136px) and (min-width: 833px)
  .search__container
    visibility visible

@media screen and (max-width: 1136px)
  >>> h1[id*='requisite'], >>> h2[id*='requisite'], >>> h3[id*='requisite'], >>> h4[id*='requisite'], >>> h5[id*='requisite'], >>> h6[id*='requisite']
    display flex

  >>> h1[id*='requisite'] + ul, >>> h2[id*='requisite'] + ul, >>> h3[id*='requisite'] + ul, >>> h4[id*='requisite'] + ul, >>> h5[id*='requisite'] + ul, >>> h6[id*='requisite'] + ul
    display block
</style>

<script>
import { findIndex, sortBy } from "lodash";
import copy from "clipboard-copy";

export default {
  props: {
    aside: {
      type: Boolean,
      default: true
    },
    tree: {
      type: Array
    }
  },
  mounted() {
    this.emitPrereqLinks();
    const headerAnchorClick = event => {
      event.target.setAttribute("data-header-anchor-text", "已拷贝!");
      copy(event.target.href);
      setTimeout(() => {
        event.target.setAttribute("data-header-anchor-text", "拷贝链接!");
      }, 1000);
      event.preventDefault();
    };
    document
      .querySelectorAll(
        'h1[id*="requisite"], h2[id*="requisite"], h3[id*="requisite"], h4[id*="requisite"], h5[id*="requisite"], h6[id*="requisite"]'
      )
      .forEach(node => {
        node.addEventListener("click", this.prereqToggle);
      });
    document
      .querySelectorAll(".content__default a.header-anchor")
      .forEach(node => {
        node.setAttribute("data-header-anchor-text", "拷贝链接");
        node.addEventListener("click", headerAnchorClick);
      });
    if (window.location.hash) {
      const elementId = document.querySelector(window.location.hash);
      if (elementId) elementId.scrollIntoView();
    }
  },
  methods: {
    emitPrereqLinks() {
      const prereq = [...document.querySelectorAll("[prereq]")].map(item => {
        const link = item.querySelector("[href]");
        return {
          href: link.getAttribute("href"),
          text: link.innerText
        };
      });
      this.$emit("prereq", prereq);
    },
    prereqToggle(e) {
      e.target.classList.toggle("prereqTitleShow");
      document.querySelectorAll("[prereq]").forEach(node => {
        node.classList.toggle("prereqLinkShow");
      });
    }
  },
  computed: {
    noAside() {
      return !this.aside;
    },
    linkPrevNext() {
      if (!this.tree) return;
      let result = {};
      const search = tree => {
        return tree.forEach(item => {
          const children = item.children;
          if (children) {
            const index = findIndex(children, ["regularPath", this.$page.path]);
            if (index >= 0 && children[index - 1]) {
              result.prev = children[index - 1];
            }
            if (index >= 0 && children[index + 1]) {
              result.next = children[index + 1];
            }
            return search(item.children);
          }
        });
      };
      search(this.tree);
      return result;
    }
  }
};
</script>
