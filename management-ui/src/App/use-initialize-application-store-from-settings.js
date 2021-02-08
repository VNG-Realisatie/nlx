// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { useContext, useEffect } from 'react'
import { useApplicationStore } from '../hooks/use-stores'
import UserContext from '../user-context'

const useInitializeApplicationStoreFromSettings = (getSettings) => {
  const { user, isReady } = useContext(UserContext)
  const applicationStore = useApplicationStore()

  useEffect(() => {
    const fetch = async () => {
      if (!isReady) return
      if (user === null) return
      if (applicationStore.isOrganizationInwaySet !== null) return

      const settings = await getSettings()

      applicationStore.update({
        isOrganizationInwaySet: !!settings.organizationInway,
      })
    }

    fetch()
  }, [user, isReady]) // eslint-disable-line react-hooks/exhaustive-deps
}

export default useInitializeApplicationStoreFromSettings
