import { css } from 'styled-components'
import placementStyling from './placementStyling'

const placements = ['top', 'bottom', 'left', 'right']
const origins = ['bottom', 'top', 'right', 'left']
const backdropOrigins = ['0% 25%', '0% -50%', '50% 0%', '-50% 0%']
const backdropBorderRadii = [
    '40% 40% 0 0',
    '0 0 30% 30%',
    '50% 0 0 50%',
    '0 50% 50% 0',
]

const popoverPlacements = placements.reduce(
    // Create an array of placement stylings, this will be joined in a string by css``
    (css, placement, i) =>
        css.concat(
            placementStyling(
                placement,
                origins[i],
                backdropOrigins[i],
                backdropBorderRadii[i],
            ),
        ),
    [],
)

const black = '#2D3240'

const tippyStyle = css`
    .tippy-iOS {
        cursor: pointer !important;
    }

    .tippy-notransition {
        transition: none !important;
    }

    .tippy-popper {
        perspective: 700px;
        z-index: 9999;
        outline: 0;
        transition-timing-function: cubic-bezier(0.165, 0.84, 0.44, 1);
        pointer-events: none;

        ${popoverPlacements}
    }

    .tippy-tooltip {
        position: relative;
        color: white;
        border-radius: 3px;
        max-width: 350px;
        text-align: center;
        will-change: transform;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        background-color: ${black};

        font-size: 12px;
        line-height: 20px;
        font-weight: 600;
        padding: 1px 8px 3px;

        &[data-animatefill] {
            overflow: hidden;
            background-color: transparent;
        }

        &[data-interactive] {
            pointer-events: auto;

            path {
                pointer-events: auto;
            }
        }

        &[data-inertia] {
            &[data-state='visible'] {
                transition-timing-function: cubic-bezier(0.53, 2, 0.36, 0.85);
            }
            &[data-state='hidden'] {
                transition-timing-function: ease;
            }
        }
    }

    .tippy-arrow,
    .tippy-roundarrow {
        position: absolute;
        width: 0;
        height: 0;
    }

    .tippy-roundarrow {
        width: 24px;
        height: 8px;
        fill: ${black};
        pointer-events: none;
    }

    .tippy-backdrop {
        position: absolute;
        will-change: transform;
        background-color: ${black};
        border-radius: 50%;
        width: calc(110% + 2rem);
        left: 50%;
        top: 50%;
        z-index: -1;
        transition: all cubic-bezier(0.46, 0.1, 0.52, 0.98);
        backface-visibility: hidden;

        &::after {
            content: '';
            float: left;
            padding-top: 100%;
        }
    }

    .tippy-backdrop + .tippy-content {
        transition-property: opacity;
        will-change: opacity;

        &[data-state='visible'] {
            opacity: 1;
        }
        &[data-state='hidden'] {
            opacity: 0;
        }
    }
`

export default tippyStyle
