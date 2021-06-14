// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

window.CP = window.CP || {};

CP.App = function (el) {
    const formEl = el.querySelector('#form')
    const requestEl = el.querySelector('#request')
    const resultEl = el.querySelector('#result')
    const crtEl = el.querySelector('#crt')
    const csrEl = el.querySelector('#csr')
    const errorEl = el.querySelector('#error')
    const downloadEl = el.querySelector('#download')

    crtEl.addEventListener('click', (event) => {
        event.currentTarget.select()
    })

    formEl.addEventListener('submit', (event) => {
        event.preventDefault()

        fetchCertificateForCSR(csrEl.value)
          .then(certificate => {
              errorEl.classList.add('d-none')
              requestEl.classList.toggle('d-none')
              resultEl.classList.toggle('d-none')

              crtEl.value = certificate

              downloadEl.setAttribute('href', `data:text/plain;charset=utf-8,${encodeURIComponent(certificate)}`)
              downloadEl.setAttribute('download', 'certificate.crt')
          })
          .catch(() => {
              errorEl.classList.remove('d-none')
          })

    })

    function fetchCertificateForCSR(csr) {
        return fetch("/api/request_certificate", {
            method: 'POST',
            body: JSON.stringify({
                csr: csr.value
            }),
            headers: {
                'Content-Type': 'application/json'
            }
        })
          .then(response => response.json())
          .then(json => json.certificate)
    }
}

window.addEventListener('DOMContentLoaded', () => {
  CP.App(document.querySelector('body'))
});
