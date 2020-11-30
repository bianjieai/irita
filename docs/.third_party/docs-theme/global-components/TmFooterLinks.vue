<template lang="pug">
  div
    .links
      .links__wrapper
        .links__container(v-if="$page.frontmatter.prev || (linkPrevNext && linkPrevNext.prev && linkPrevNext.prev.frontmatter && linkPrevNext.prev.frontmatter.order !== false)")
          //- .links__label Previous
          router-link.links__item.links__item__left(:to="$page.frontmatter.prev || linkPrevNext.prev.regularPath")
            .links__item__icon
              svg(width="100%" height="100%" viewBox="0 0 44 32" fill="none" xmlns="http://www.w3.org/2000/svg")
                path(d="M43 17C43.5523 17 44 16.5523 44 16C44 15.4477 43.5523 15 43 15L43 17ZM1.5 16L0.792896 15.2929L0.085789 16L0.792896 16.7071L1.5 16ZM15.7929 31.7071C16.1834 32.0976 16.8166 32.0976 17.2071 31.7071C17.5976 31.3166 17.5976 30.6834 17.2071 30.2929L15.7929 31.7071ZM17.2071 1.70711C17.5976 1.31658 17.5976 0.683419 17.2071 0.292895C16.8166 -0.0976288 16.1834 -0.0976289 15.7929 0.292895L17.2071 1.70711ZM43 15L1.5 15L1.5 17L43 17L43 15ZM0.792896 16.7071L15.7929 31.7071L17.2071 30.2929L2.20711 15.2929L0.792896 16.7071ZM15.7929 0.292895L0.792896 15.2929L2.20711 16.7071L17.2071 1.70711L15.7929 0.292895Z" fill="#0E2125" fill-opacity="0.26")
            div
              .links__item__title {{$page.frontmatter.prev || linkPrevNext.prev.title}}
              .links__item__desc(v-if="linkPrevNext.prev.frontmatter.description" v-html="shorten(linkPrevNext.prev.frontmatter.description)")
      .links__wrapper
        .links__container(v-if="$page.frontmatter.next || (linkPrevNext && linkPrevNext.next && linkPrevNext.next.frontmatter && linkPrevNext.next.frontmatter.order !== false)")
          //- .links__label Up next
          router-link.links__item.links__item__right(:to="$page.frontmatter.next || linkPrevNext.next.regularPath")
            div
              .links__item__title {{$page.frontmatter.next || linkPrevNext.next.title}}
              .links__item__desc(v-if="linkPrevNext.next.frontmatter.description" v-html="shorten(linkPrevNext.next.frontmatter.description)")
            .links__item__icon
              svg(width="100%" height="100%" viewBox="0 0 44 32" fill="none" xmlns="http://www.w3.org/2000/svg")
                path(d="M0.999994 17C0.447709 17 -6.34082e-06 16.5523 -6.43738e-06 16C-6.53395e-06 15.4477 0.447709 15 0.999993 15L0.999994 17ZM42.5 16L43.2071 15.2929L43.9142 16L43.2071 16.7071L42.5 16ZM28.2071 31.7071C27.8166 32.0976 27.1834 32.0976 26.7929 31.7071C26.4024 31.3166 26.4024 30.6834 26.7929 30.2929L28.2071 31.7071ZM26.7929 1.70711C26.4024 1.31658 26.4024 0.683419 26.7929 0.292895C27.1834 -0.0976288 27.8166 -0.0976289 28.2071 0.292895L26.7929 1.70711ZM0.999993 15L42.5 15L42.5 17L0.999994 17L0.999993 15ZM43.2071 16.7071L28.2071 31.7071L26.7929 30.2929L41.7929 15.2929L43.2071 16.7071ZM28.2071 0.292895L43.2071 15.2929L41.7929 16.7071L26.7929 1.70711L28.2071 0.292895Z" fill="#0E2125" fill-opacity="0.26")
</template>

<style lang="stylus" scoped>
.links
  display flex

  &__wrapper
    display flex
    width 100%
    margin-bottom 2rem

    &:first-child
      margin-right 2rem

  &__container
    width 100%
    align-items stretch
    display flex
    flex-direction column
    // background red

  &__item
    margin-top 1rem
    padding 2rem
    box-shadow 0px 2px 4px rgba(22, 25, 49, 0.05), 0px 0px 1px rgba(22, 25, 49, 0.2), 0px 0.5px 0px rgba(22, 25, 49, 0.05)
    border-radius 0.5rem
    display grid
    grid-auto-flow column
    flex-grow 1
    gap 2rem
    overflow-x hidden

    &__left
      grid-template-columns 44px auto

    &__right
      grid-template-columns auto 44px

    &__icon
      display flex
      align-items center

    &__title
      margin-top 5px
      font-weight 500
      font-size 1.25rem

    &__desc
      color rgba(22, 25, 49, 0.65)
      margin-top 0.5rem
      font-size 0.875rem
      line-height 20px

  &__label
    color rgba(22, 25, 49, 0.9)
    text-transform uppercase
    font-size 0.75rem
    letter-spacing 0.2rem

@media screen and (max-width: 1280px)
  .links
    flex-direction column-reverse
</style>

<script>
import { findIndex, find } from "lodash";

export default {
  props: ["tree"],
  methods: {
    shorten(string) {
      let str = string.split(" ");
      str =
        str.length > 20 ? str.slice(0, 20).join(" ") + "..." : str.join(" ");
      return this.md(str);
    }
  },
  computed: {
    linkPrevNext() {
      if (!this.tree) return;
      let result = {};
      const search = tree => {
        return tree.forEach((item, i) => {
          const children = item.children;
          if (children) {
            const index = findIndex(children, ["regularPath", this.$page.path]);
            if (index >= 0 && children[index - 1]) {
              result.prev = children[index - 1];
            }
            if (index >= 0 && children[index + 1]) {
              result.next = children[index + 1];
            } else if (index >= 0 && tree[i + 1] && tree[i + 1].children) {
              result.next = find(tree[i + 1].children, x => {
                return x.frontmatter && x.frontmatter.order !== false;
              });
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