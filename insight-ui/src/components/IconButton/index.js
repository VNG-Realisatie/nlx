import styled from 'styled-components'
import { bool, string } from 'prop-types'

const IconButton = styled.button`
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;

    width: 48px;
    height: 48px;
    color: #517FFF;

    user-select: none;
    background: none;
    border: none;
    border-radius: ${p => p.rounded === 'true' ? '50%' : 0};
    overflow: hidden;

    &:not([disabled]) {
        cursor: pointer;
    }

    &[disabled] {
        color: #CAD0E0;
        cursor: default;
        pointer-events: none;
    }

    /* Enable a transparent layer over the original background color of the button (for :hover, :focus, :active)... */
    &:before {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
    }

    /* ...and content stays on top of transparent layer */
    > * {
        position: relative;
    }

    &:hover {
        &:before {
            background-color: rgba(81,127,255,0.04);
        }
    }
    &:active {
        &:before {
            background-color: rgba(81,127,255,0.08);
        }
    }
`

IconButton.propTypes = {
  rounded: string,
  disabled: bool,
}

IconButton.defaultProps = {
  rounded: 'true',
  disabled: false,
}

export default IconButton
