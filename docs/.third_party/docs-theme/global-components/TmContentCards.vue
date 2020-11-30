<template lang="pug">
  div
    .container
      .title Contents
      router-link(v-for="card in cards" :to="card.path" tag="a").card
        .card__title {{card.title}}
        .card__description(v-if="card.frontmatter.description") {{card.frontmatter.description}}
        svg(width="12" height="20" viewBox="0 0 12 20" fill="none" xmlns="http://www.w3.org/2000/svg").card__icon
          path(d="M1.5 1.75L9.75 10L1.5 18.25" stroke="#A2A3AD" stroke-width="2" stroke-linecap="round")
</template>

<style lang="stylus" scoped>
.container
  margin-top 3rem

.title
  font-size 1.5rem
  font-weight 600
  margin-bottom 2rem

.card
  border 1px solid rgba(140, 145, 177, 0.32)
  border-radius 8px
  margin-bottom 1.5rem
  padding 1.5rem 2rem
  position relative
  display block
  color inherit

  &__title
    color #161931
    font-size 1.25rem
    font-weight 600

  &__description
    margin-top .75rem
    font-size .875rem
    line-height 20px

  &__icon
    top 1.5rem
    right 1.5rem
    position absolute
</style>

<script>
export default {
  computed: {
    cards() {
      return this.$site.pages
        .filter(page => page.path.match(this.$page.path))
        .filter(page => page.path != this.$page.path);
    }
  }
};
</script>