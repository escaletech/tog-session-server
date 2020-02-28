const splitFlags = (flags, state) =>
  flags && flags.split(',').reduce((all, flag) => ({ ...all, [flag]: state }), {})

const parseOptions = ({ enable, disable }) => {
  return {
    flags: {
      ...splitFlags(enable, true),
      ...splitFlags(disable, false)
    }
  }
}

module.exports = {
  parseOptions
}
