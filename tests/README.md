const dependencies = {
  react: '17.0.2',
  'react-native': '0.68.2',
  '@react-native-community/netinfo': '6.1.0',
  'react-native-navigation': '5.5.0',
  'react-native-vector-icons': '9.1.0',
  'react-native-safe-area-context': '3.2.0',
  'lodash.debounce': '4.1.3',
  'react-redux': '7.2.5',
  'redux': '4.1.2',
  'react-redux-firebase': '2.1.9',
  'firebase': '9.6.11',
  'realm': '10.24.0',
};

const devDependencies = {
  'jest': '27.4.4',
  'jest-react-native': '18.0.1',
  'babel-jest': '27.4.6',
  'react-test-renderer': '17.0.2',
  'enzyme': '3.11.0',
  'enzyme-adapter-react-16': '1.15.6',
  'babel-preset-react-native': '4.0.1',
  'babel-preset-es2020': '7.2.0',
};

const peerDependencies = {
  'react-native-screens': '3.10.1',
  'react-native-safe-area-context': '3.2.0',
};

module.exports = {
  name: 'mobile-app-react-native',
  version: '1.0.0',
  main: 'index.js',
  scripts: {
    start: 'react-native start',
    test: 'jest',
  },
  dependencies,
  devDependencies,
  peerDependencies,
};