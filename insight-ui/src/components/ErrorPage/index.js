import React from 'react'
import { StyledCard } from './index.styles'

const ErrorPage = ({ title, children, ...props}) =>
  <StyledCard {...props}>
    <h1>{title}</h1>
    {children}
  </StyledCard>

export default ErrorPage
