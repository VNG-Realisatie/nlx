// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, instanceOf, string } from 'prop-types'
import { observer } from 'mobx-react'
import Table from '../../../../../../../../components/Table'
import { OutwayName, Outways } from '../../components/index.styles'
import AccessState from '../../components/AccessState'
import OutwayModel from '../../../../../../../../stores/models/OutwayModel'
import DirectoryServiceModel from '../../../../../../../../stores/models/DirectoryServiceModel'

const Row = ({ publicKeyFingerprint, outways, service }) => {
  const { accessRequest, accessProof } =
    service.getAccessStateFor(publicKeyFingerprint)

  return (
    <Table.Tr>
      <Table.Td>
        <Outways>
          {outways.map((outway) => (
            <OutwayName key={outway.name}>{outway.name}</OutwayName>
          ))}
        </Outways>

        <AccessState accessRequest={accessRequest} accessProof={accessProof} />
      </Table.Td>
    </Table.Tr>
  )
}

Row.propTypes = {
  publicKeyFingerprint: string,
  outways: arrayOf(instanceOf(OutwayModel)),
  service: instanceOf(DirectoryServiceModel),
}

export default observer(Row)
