// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { ReactComponent as IconBox } from './box.svg'
import { ReactComponent as IconBuilding } from './building.svg'
import { ReactComponent as IconCheck } from './check.svg'
import { ReactComponent as IconExternalLink } from './external-link.svg'
import { ReactComponent as IconFileCopy } from './file-copy-2.svg'
import { ReactComponent as IconGitlab } from './gitlab.svg'
import { ReactComponent as IconGroup } from './group-2.svg'
import { ReactComponent as IconHammer } from './hammer.svg'
import { ReactComponent as IconHome } from './home.svg'
import { ReactComponent as IconInfo } from './info.svg'
import { ReactComponent as IconLightning } from './flashlight.svg'
import { ReactComponent as IconMail } from './mail.svg'
import { ReactComponent as IconOpenArm } from './open-arm.svg'
import { ReactComponent as IconPlug } from './plug.svg'
import { ReactComponent as IconRecycle } from './recycle.svg'
import { ReactComponent as IconShieldCheck } from './shield-check.svg'
import { ReactComponent as IconSpy } from './spy.svg'
import { ReactComponent as IconTools } from './tools.svg'

const icons = {
  box: IconBox,
  building: IconBuilding,
  check: IconCheck,
  externalLInk: IconExternalLink,
  fileCopy: IconFileCopy,
  gitlab: IconGitlab,
  group: IconGroup,
  hammer: IconHammer,
  home: IconHome,
  info: IconInfo,
  lightning: IconLightning,
  mail: IconMail,
  openArm: IconOpenArm,
  plug: IconPlug,
  recycle: IconRecycle,
  shieldCheck: IconShieldCheck,
  spy: IconSpy,
  tools: IconTools,
}

export {
  IconBox,
  IconBuilding,
  IconCheck,
  IconExternalLink,
  IconFileCopy,
  IconGitlab,
  IconGroup,
  IconHammer,
  IconHome,
  IconInfo,
  IconLightning,
  IconMail,
  IconOpenArm,
  IconPlug,
  IconRecycle,
  IconShieldCheck,
  IconSpy,
  IconTools,
}

export const getIcon = (icon) => {
  if (Object.keys(icons).includes(icon)) {
    // Above check makes it input-safe, so ignore eslint
    // eslint-disable-next-line security/detect-object-injection
    return icons[icon]
  }

  if (!process.env.NEXT_PUBLIC_PRODUCTION) {
    console.warn(`Icon "${icon}" not found`)
  }

  return icons.tools
}
