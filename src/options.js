const splitFlags = (flags, state) =>
  flags && flags.split(',').reduce((all, flag) => ({ ...all, [flag]: state }), {})

const parseOptions = ({ enable, disable, duration }) => {
  return {
    flags: {
      ...splitFlags(enable, true),
      ...splitFlags(disable, false)
    },
    duration: Number(duration)
  }
}

module.exports = {
  parseOptions
}
