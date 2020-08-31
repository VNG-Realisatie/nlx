// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import fs from 'fs'
import path from 'path'

export default async function ({ testController, requests, directory, fileName }) {
  const filePath = path.join(directory, fileName)

  const result = {
    fixture: testController.testRun.test.testFile.currentFixture.name,
    testName: testController.testRun.test.name,
    requests: requests
  }

  await fs.promises.writeFile(filePath, JSON.stringify(result, null, 4))
}
