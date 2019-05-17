// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import styled from 'styled-components'

const Small = styled.div`
    font-family: ${p => p.theme.font.family.main};
    font-size: ${p => p.theme.font.size.small};
    line-height: ${p => p.theme.font.lineHeight.small};
    color: ${p => p.theme.color.grey[60]};
`

export default Small
