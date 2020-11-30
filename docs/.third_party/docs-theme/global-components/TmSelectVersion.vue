<template lang="pug">
  div
    .select(v-if="versions.length > 1")
      select(@input="select")
        option(v-for="item in versions" :value="item") {{item}}
</template>

<style lang="stylus" scoped>
select
  border none
  background none
  text-transform uppercase
  letter-spacing 0.02em
  font-weight 500
  font-size 0.875rem

.select
  border 2px solid rgba(140, 145, 177, 0.32)
  padding 0.25rem 0.5rem
  border-radius 6px
  background-color transparent
</style>

<script>
export default {
  computed: {
    versions() {
      return this.$themeConfig.versions;
    }
  },
  methods: {
    select(e) {
      // this.$router.push(
      //   this.$page.path.replace(this.$localeConfig.path, e.target.value)
      // );
    }
  }
};
</script>