import { css } from 'styled-components'
import {
    arrow,
    arrowPosition,
    arrowMargin,
    roundarrowTransform,
    arrowTransformOrigin,
    roundarrowTransformOrigin,
} from './helpers/arrow'
import {
    backdropTransformEnter,
    backdropTransformLeave,
} from './helpers/backdrop'
import { enterTransform, leaveTransform } from './helpers/transform'

export default function placementStyling(
    placement,
    origin,
    backdropOrigin,
    backdropBorderRadius,
) {
    return css`
        &[x-placement^="${placement}"] {
            .tippy-backdrop {
                border-radius: ${backdropBorderRadius};
            }

            .tippy-roundarrow {
                ${arrowPosition(placement, '8px', true)};
                ${roundarrowTransformOrigin(placement)};

                svg {
                    position: absolute;
                    left: 0;
                    ${roundarrowTransform(placement)};
                }
            }

            .tippy-arrow {
                ${(p) => arrow(placement, '5px', 'black')};
                ${arrowPosition(placement, '5px', false)};
                ${arrowMargin(placement)};
                ${arrowTransformOrigin(placement)};
            }

            .tippy-backdrop {
                transform-origin: ${backdropOrigin};

                &[data-state='visible'] {
                    ${backdropTransformEnter(placement)};
                }

                &[data-state='hidden'] {
                    ${backdropTransformLeave(placement)};
                    opacity: 0;
                }
            }

            [data-animation='shift-toward'] {
                &[data-state='visible'] {
                    ${enterTransform(placement, 'fade')}
                }
                &[data-state='hidden'] {
                    opacity: 0;
                    ${leaveTransform(placement, 'shift-toward')}
                }
            }

            [data-animation='perspective'] {
                transform-origin: ${origin};

                &[data-state='visible'] {
                    ${enterTransform(placement, 'perspective')}
                }
                &[data-state='hidden'] {
                    opacity: 0;
                    ${leaveTransform(placement, 'perspective')}
                }
            }

            [data-animation='fade'] {
                &[data-state='visible'] {
                    ${enterTransform(placement, 'fade')}
                }
                &[data-state='hidden'] {
                    opacity: 0;
                    ${enterTransform(placement, 'fade')}
                }
            }

            [data-animation='shift-away'] {
                &[data-state='visible'] {
                    ${enterTransform(placement, 'fade')}
                }
                &[data-state='hidden'] {
                    opacity: 0;
                    ${leaveTransform(placement, 'shift-away')}
                }
            }

            [data-animation='scale'] {
                &[data-state='visible'] {
                    ${enterTransform(placement, 'scale')}
                }
                &[data-state='hidden'] {
                    opacity: 0;
                    ${leaveTransform(placement, 'scale')}
                }
            }
        }
    `
}
