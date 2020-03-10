const splitFlags = (flags, state) =>
  flags && flags.split(',').reduce((all, flag) => ({ ...all, [flag]: state }), {})

const parseOptions = ({ enable, disable, duration }) => {
  return {
    flags: {
      ...splitFlags(enable, true),
      ...splitFlags(disable, false)
    },
    duration: parseInt(duration)
  }
}

module.exports = {
  parseOptions
}
