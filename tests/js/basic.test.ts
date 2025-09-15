describe('Basic Tests', () => {
  it('should pass basic test', () => {
    expect(true).toBe(true);
  });
  
  it('should have environment', () => {
    expect(process.env).toBeDefined();
  });
});
