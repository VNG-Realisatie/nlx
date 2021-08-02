// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { createMedia } from '@artsy/fresnel'
import { defaultTheme } from '@commonground/design-system'

const fresnel = createMedia({
  breakpoints: defaultTheme.breakpoints,
})

export const mediaStyles = fresnel.createMediaStyle()
export const { Media, MediaContextProvider: MediaProvider } = fresnel
