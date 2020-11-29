<template lang="pug">
  div
    .container(v-if="$themeConfig && $themeConfig.gutter")
      .wrapper
        .title(v-if="$themeConfig.gutter.title") {{$themeConfig.gutter.title}}
        .links
          a(:href="$themeConfig.gutter.chat.url" target="_blank" rel="noreferrer noopener" :style="{'--bg': $themeConfig.gutter.chat.bg}").links__item.links__item__chat
            tm-logo-chat.links__item__logo(style="fill: white; width: 80px; height: 80px; padding: 10px;")
            div(v-html="md($themeConfig.gutter.chat.title)").links__item__title
            div(v-html="md($themeConfig.gutter.chat.text)").links__item__text
          a(:href="$themeConfig.gutter.forum.url" target="_blank" rel="noreferrer noopener" :style="{'--bg': $themeConfig.gutter.forum.bg}").links__item.links__item__forum
            component(:is="`tm-logo-${$themeConfig.gutter.forum.logo}`" style="fill: white; width: 100px; height: 100px;").links__item__logo
            div(v-html="md($themeConfig.gutter.forum.title)").links__item__title
            div(v-html="md($themeConfig.gutter.forum.text)").links__item__text
          a(:href="editLink" target="_blank" rel="noreferrer noopener").links__item.links__item__regular
            tm-icon-paper-pen(fill="var(--accent-color").links__item__logo
            div(v-html="md($themeConfig.gutter.github.title)").links__item__title
            div(v-html="md($themeConfig.gutter.github.text)" style="color: rgba(22, 25, 49, 0.65)").links__item__text
      .newsletter(v-if="$themeConfig.label == 'sdk'")
        .newsletter__image
          .newsletter__image__item(v-for="(item, index) in range(15)" :class="[`letter__${index}`]")
            image-letter
        .newsletter__form
          tm-newsletter-form
</template>

<style lang="stylus" scoped>
.newsletter
  box-shadow 0px 2px 4px rgba(22, 25, 49, 0.05), 0px 0px 1px rgba(22, 25, 49, 0.2), 0px 0.5px 0px rgba(22, 25, 49, 0.05)
  margin-top 4rem
  margin-bottom 1rem
  overflow hidden
  min-height 200px
  border-radius 0.5rem
  position relative
  display flex
  align-items center
  justify-content flex-end

  &:hover &__image__item
    opacity 0.25
    transform scale(0.98)

  &:hover &__image__item.letter__10
    transform scale(1.5) rotate(25deg)
    opacity 1

  &__image
    transform translate(-125px, -35px) rotate(-25deg)
    display grid
    width 300px
    gap 1rem
    grid-template-columns repeat(4, 1fr)
    position absolute
    left 0
    top 0

    &__item
      opacity 0.5
      transition all 0.5s

      &.letter__10
        opacity 1
        background white

  &__form
    margin-left 250px
    padding 30px
    width 100%

/deep/
  a[href]
    color var(--accent-color)

  strong
    font-weight 600

strong
  font-weight 500

.container
  background var(--sidebar-bg)

.wrapper
  max-width calc(1400px - var(--sidebar-width))

.title
  font-size 2rem
  color #161931
  padding 1.5rem 0
  font-weight 600

.links
  display grid
  gap 2rem
  grid-template-columns repeat(auto-fit, minmax(250px, 1fr))

  &__item
    display flex
    flex-direction column
    align-items center
    box-shadow 0px 2px 4px rgba(22, 25, 49, 0.05), 0px 0px 1px rgba(22, 25, 49, 0.2), 0px 0.5px 0px rgba(22, 25, 49, 0.05)
    text-align center
    border-radius 0.5rem
    padding 2rem
    line-height 20px
    background var(--bg)
    transition box-shadow .25s

    &__text
      font-size .875rem
      line-height 20px
      letter-spacing 0.01em

    &:hover
      box-shadow 0px 12px 24px rgba(22, 25, 49, 0.07), 0px 4px 8px rgba(22, 25, 49, 0.05), 0px 1px 0px rgba(22, 25, 49, 0.05)

    &__title
      margin-top 1.5rem
      margin-bottom 1rem
      font-weight 600

a.links__item
  color white

a.links__item__regular
  color #161931
  background rgba(176, 180, 207, 0.09)

@media screen and (max-width: 832px)
  .newsletter
    height initial

    &__form
      margin-top 250px
      margin-left initial
      grid-template-columns 1fr
</style>

<script>
import { range } from "lodash";

const endingSlashRE = /\/$/;
const outboundRE = /^[a-z]+:/i;

export default {
  computed: {
    editLink() {
      if (this.$page.frontmatter.editLink === false) {
        return;
      }
      const {
        repo,
        editLinks,
        docsDir = "",
        docsBranch = "master",
        docsRepo = repo
      } = this.$site.themeConfig;
      if (docsRepo && editLinks && this.$page.relativePath) {
        return this.createEditLink(
          repo,
          docsRepo,
          docsDir,
          docsBranch,
          this.$page.relativePath
        );
      }
    },
    editLinkText() {
      return (
        this.$themeLocaleConfig.editLinkText ||
        this.$site.themeConfig.editLinkText ||
        `Edit this page`
      );
    }
  },
  methods: {
    createEditLink(repo, docsRepo, docsDir, docsBranch, path) {
      const bitbucket = /bitbucket.org/;
      if (bitbucket.test(repo)) {
        const base = outboundRE.test(docsRepo) ? docsRepo : repo;
        return (
          base.replace(endingSlashRE, "") +
          `/src` +
          `/${docsBranch}/` +
          (docsDir ? docsDir.replace(endingSlashRE, "") + "/" : "") +
          path +
          `?mode=edit&spa=0&at=${docsBranch}&fileviewer=file-view-default`
        );
      }
      const base = outboundRE.test(docsRepo)
        ? docsRepo
        : `https://github.com/${docsRepo}`;
      return (
        base.replace(endingSlashRE, "") +
        `/edit` +
        `/${docsBranch}/` +
        (docsDir ? docsDir.replace(endingSlashRE, "") + "/" : "") +
        path
      );
    },
    range
  }
};
</script>