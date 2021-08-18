// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { Container } from './index.styles'

const ErrorMessage = () => (
  <Container>
    <h1>Kon geen informatie ophalen</h1>
    <p>
      Probeer het later nog eens.
      <br />
      Excuus voor het ongemak.
    </p>
  </Container>
)

export default ErrorMessage
