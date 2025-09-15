module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  roots: ['<rootDir>/tests/js'],
  testMatch: ['**/*.test.ts'],
  collectCoverageFrom: [
    'src/js/**/*.ts',
    '!src/js/**/*.d.ts',
    '!src/js/index.ts',
  ],
  coverageDirectory: 'coverage',
  coverageReporters: ['text', 'lcov', 'html'],
};