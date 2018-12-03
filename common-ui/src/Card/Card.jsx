import styled from 'styled-components'

const Card = styled.div`
    background-color: ${p => p.theme.color.white};
    border-radius: ${p => p.theme.radius.small};
    border: 1px solid ${p => p.theme.color.grey[30]};
`

export default Card
export { default as CardContent } from './CardContent'
