import jwt from "jsonwebtoken";

const generateToken = userId => {
  return jwt.sign({ userId }, "my-jwt-secret", { expiresIn: "7 days" });
};

export { generateToken as default };
