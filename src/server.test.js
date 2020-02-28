const server = require('./server')

const mockSession = jest.fn()

jest.mock('tog-node', () => ({
  SessionClient: jest.fn().mockImplementation(() => ({ session: mockSession }))
}))

const any = expect.anything()
const defaultExpiration = 60
const mockSessionPayload = { id: 'abc123' }
const serializedPayload = JSON.stringify(mockSessionPayload)

beforeEach(() => {
  mockSession.mockRestore()
  mockSession.mockResolvedValue(mockSessionPayload)
})

describe('retrieve session', () => {
  describe('with path prefix', () => {
    const app = server({ pathPrefix: '/sessions', defaultExpiration })

    afterAll(() => app.close())

    test('accepts requests with prefix', async () => {
      const res = await app.inject('/sessions/some-ns/abc123')

      expect(res.statusCode).toBe(200)
      expect(res.payload).toEqual(serializedPayload)

      expect(mockSession).toHaveBeenCalledWith(
        'some-ns', 'abc123', expect.objectContaining({ duration: defaultExpiration }))
    })
  })

  describe('with default configuration', () => {
    const app = server({ pathPrefix: '', defaultExpiration })

    afterAll(() => app.close())

    test('returns session information', async () => {
      const res = await app.inject('/some-ns/abc123')

      expect(res.statusCode).toBe(200)
      expect(res.payload).toEqual(serializedPayload)

      expect(mockSession).toHaveBeenCalledWith(
        'some-ns', 'abc123', expect.objectContaining({ duration: defaultExpiration, flags: {} }))
    })

    test('accepts enabled flags', async () => {
      await app.inject('/some-ns/abc123?enable=one,two')

      expect(mockSession).toHaveBeenCalledWith(any, any,
        expect.objectContaining({
          flags: { one: true, two: true }
        }))
    })

    test('accepts disabled flags', async () => {
      await app.inject('/some-ns/abc123?disable=one,two')

      expect(mockSession).toHaveBeenCalledWith(any, any,
        expect.objectContaining({
          flags: { one: false, two: false }
        }))
    })
  })
})
