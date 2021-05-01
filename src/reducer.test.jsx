import reducer from "./reducer";

describe("reducer", () => {
  it("should return the initial state when receiving an undefined action", () => {
    const expected = {};
    const store = undefined;
    const action = "UNKNOWN";
    expect(reducer(store, action)).toEqual(expected);
  });
});
