import styled from 'styled-components'

export const StyledMain = styled.main`
  display: flex;
  align-items: flex-start;
  height: 100%;
`

export const StyledContent = styled.div`
  flex: 1;
  padding: ${(p) => p.theme.tokens.spacing07} ${(p) => p.theme.tokens.spacing09};
`

export const StyledPageTitle = styled.h1`
  margin-bottom: ${(p) => p.theme.tokens.spacing01};
`

export const StyledPageDescription = styled.p`
  margin-bottom: ${(p) => p.theme.tokens.spacing07};
`

export const StyledPageHeader = styled.div`
  display: flex;
  justify-content: space-between;
`
