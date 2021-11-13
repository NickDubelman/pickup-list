interface QueryInput {
  query: string
  variables?: any
}

export async function graphqlQuery<T>(
  fetch: any,
  { query, variables }: QueryInput
): Promise<T> {
  const resp = await fetch('http://localhost:8080/query', {
    method: 'POST',
    headers: { 'content-type': 'application/json' },
    body: JSON.stringify({ query, variables })
  })

  const { data, errors } = await resp.json()
  if (errors && errors.length > 0) {
    throw new Error(errors[0].message)
  }

  return data
}
