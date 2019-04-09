import React from 'react'
import { StyledCard } from './index.styles'

const Home = () =>
  <StyledCard>
    <p>
      View logs by selecting an organization on the left.
      You can only view logs by disclosing the required IRMA attributes.
    </p>

    <p className="text-muted">
      Read more about IRMA and what it does <a href="https://privacybydesign.foundation/irma/" target="_blank" rel="noopener noreferrer">here</a>.
    </p>
  </StyledCard>

export default Home
