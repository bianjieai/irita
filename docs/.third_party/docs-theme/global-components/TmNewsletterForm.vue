<template lang="pug">
  form(action="https://zcs1.maillist-manage.com/campaigns/weboptin.zc" method="POST" target="_blank" rel="noreferrer noopener").newsletter__form__component__internal
    .newsletter__desc
      .newsletter__desc__h1
        label(for="newsletter-email-input") Get SDK updates
      .newsletter__desc__p You’ll be notified with Cosmos updates, news and announcements.
    .newsletter__input
      .newsletter__input__input
        input(type="email" required name="CONTACT_EMAIL" id="newsletter-email-input" placeholder="Your email").newsletter__input__input__el
        button(type="submit").newsletter__input__input__button Sign up
        input#submitType(type="hidden" name="submitType" value="optinCustomView")
        input#lD(type="hidden" name="lD" value="16352f88325b24db")
        input#emailReportId(type="hidden" name="emailReportId" value="")
        input#cmpZuid(type="hidden" name="zx" value="129a50c11")
        input(type="hidden" name="zcvers" value="2.0")
        input#allCheckedListIds(type="hidden" name="oldListIds" value="")
        input#mode(type="hidden" name="mode" value="OptinCreateView")
        input#zcld(type="hidden" name="zcld" value="16352f88325b24db")
        input#document_domain(type="hidden" value="zoho.com")
        input#zc_Url(type="hidden" value="zcs1.maillist-manage.com")
        input#new_optin_response_in(type="hidden" value="1")
        input#duplicate_optin_response_in(type="hidden" value="1")
        input#zc_formIx(type="hidden" name="zc_formIx" value="4ef47fbb86ab66681d4cb3275283cf70ab16e7ccaa8dd327")
        input#scriptless(type="hidden" name="scriptless" value="yes")
      .newsletter__input__p You can unsubscribe at any time.
</template>

<style lang="stylus" scoped>
.newsletter__form__component__internal
  display grid
  grid-template-columns repeat(auto-fit, minmax(250px, 1fr))
  width 100%
  align-items flex-end
  background-color white

.newsletter

  &__desc
    padding-right 1.5rem

    &__h1
      color #161931
      font-size 1.25rem
      font-weight 500
      margin-bottom 1rem
    
    &__p
      font-size .875rem
      line-height 20px
      color rgba(22, 25, 49, 0.9)

  &__input
    padding-top 1rem

    &__input
      border 1px solid rgba(140, 145, 177, 0.32)
      display grid
      grid-template-columns 1fr auto
      align-items center
      padding .5rem
      border-radius .25rem

      &:focus-within
        border-color #66A1FF
        
      &__button
        padding .5rem
        text-transform uppercase
        color var(--accent-color)
        font-weight 500
        cursor pointer
        background none
        border none

      &__el
        border none
        width 100%
        outline none
        padding .5rem
        background none
  
    &__p
      margin-top .5rem
      color rgba(22, 25, 49, 0.65)
      font-size .8125rem

@media screen and (max-width: 832px)
  .newsletter__form__component__internal
    display grid
    grid-template-columns repeat(auto-fit, minmax(300px, 1fr))

  .newsletter

    &__desc
      padding-right 0
      padding-bottom 1rem

@media screen and (max-width: 480px)
  .newsletter__form__component__internal
      grid-template-columns repeat(auto-fit, minmax(200px, 1fr))
</style>