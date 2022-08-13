document.addEventListener('DOMContentLoaded', () => {
  // Functions to open and close a modal
  function openModal($el) {
    $el.classList.add('is-active')
  }

  function closeModal($el) {
    closeScroll()
    $el.classList.remove('is-active')
  }

  function closeAllModals() {
    ;(document.querySelectorAll('.modal') || []).forEach(($modal) => {
      closeModal($modal)
    })
  }

  // Add a click event on buttons to open a specific modal
  ;(document.querySelectorAll('.js-modal-trigger') || []).forEach(
    ($trigger) => {
      const modal = $trigger.dataset.target
      const $target = document.getElementById(modal)

      $trigger.addEventListener('click', () => {
        openModal($target)
      })
    },
  )

  // Add a click event on various child elements to close the parent modal
  ;(
    document.querySelectorAll(
      '.modal-background, .modal-close, .modal-card-head .delete, .modal-card-foot .button',
    ) || []
  ).forEach(($close) => {
    const $target = $close.closest('.modal')

    $close.addEventListener('click', () => {
      closeModal($target)
    })
  })

  // Add a keyboard event to close all modals
  document.addEventListener('keydown', (event) => {
    const e = event || window.event

    if (e.keyCode === 27) {
      // Escape key
      closeAllModals()
    }
  })
})

function lockScroll() {
  var scrollPosition = [
    self.pageXOffset ||
      document.documentElement.scrollLeft ||
      document.body.scrollLeft,
    self.pageYOffset ||
      document.documentElement.scrollTop ||
      document.body.scrollTop,
  ]
  var html = jQuery('html') // it would make more sense to apply this to body, but IE7 won't have that
  html.data('scroll-position', scrollPosition)
  html.data('previous-overflow', html.css('overflow'))
  html.css('overflow', 'hidden')
  window.scrollTo(scrollPosition[0], scrollPosition[1])
}

function closeScroll() {
  var html = jQuery('html')
  var scrollPosition = html.data('scroll-position')
  html.css('overflow', html.data('previous-overflow'))
  window.scrollTo(scrollPosition[0], scrollPosition[1])
}
