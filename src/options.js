const splitFlags = (flags, state) =>
  flags && flags.split(',').reduce((all, flag) => ({ ...all, [flag]: state }), {})

const parseOptions = ({ experiment, enable, disable }) => {
  return {
    experiment: experiment || undefined,
    flags: {
      ...splitFlags(enable, true),
      ...splitFlags(disable, false)
    }
  }
}

module.exports = {
  parseOptions
}
