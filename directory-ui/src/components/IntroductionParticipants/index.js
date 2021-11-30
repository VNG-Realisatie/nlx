// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { Container } from '../Grid'
import { SectionIntro } from '../Section'
import { Section, Content } from './index.styles'

const Introduction = () => (
  <Section omitArrow>
    <Container>
      <SectionIntro>
        <Content>
          <h1>Directory Deelnemers</h1>
          <p>
            In dit overzicht vindt u alle deelnemers aan dit NLX ecosysteem
            (demo, pre-productie en productie).
          </p>
        </Content>
      </SectionIntro>
    </Container>
  </Section>
)

export default Introduction
