import { sign, verify } from "jsonwebtoken";
import { z } from "zod";
import { userSchema } from "~/adapter/in/dto/user";
import { isNullish } from "~/infrastructure/utils/helpers.ts";
import { UserRole } from "~/schema/types.generated";


/**
 * Represents a user session.
 */
export interface Session {
  /**
   * The ID of the user.
   */
  id: string;
  /**
   * The name of the user.
   */
  username: string;
  /**
   * The role of the user.
   */
  role: UserRole;
  /**
   * The session id.
   */
  sessionId: string;
  /**
   * The last login time.
   */
  lastLogin: string;
}

/**
 * A service that handles user sessions.
 */
export class SessionService {
  /**
   * The private key which is used to sign the session.
   */
  readonly #privateKey: string;

  /**
   * The public key which is used to verify the session.
   */
  readonly #publicKey: string;

  constructor(privateKey: string, publicKey: string) {
    this.#privateKey = privateKey;
    this.#publicKey = publicKey;
  }

  /**
   * Logins the user.
   *
   * @param request The request object.
   * @param user The user to login.
   */
  createSession(request: Request, user: z.infer<typeof userSchema>) {
    if (isNullish(user.sessionId)) {
      throw new Error("Invalid session");
    }

    const jwt = this.sign({
      id: user.id,
      username: user.username,
      role: user.role,
      sessionId: user.sessionId,
      lastLogin: user.lastLogin,
    });

    // Set the cookie
    request.cookieStore?.set({
      domain: null,
      path: "/",
      name: "jwt",
      value: jwt,
      httpOnly: true,
      secure: process.env.SECURE_COOKIES === "true",
      expires: new Date(Date.now() + 1000 * 60 * 60 * 24 * 7),
    });
  }

  /**
   * Logs out the user.
   *
   * @param request The request object.
   */
  destroySession(request: Request) {
    // Delete the cookie
    request.cookieStore?.delete("jwt");
  }

  /**
   * Signs the session data.
   *
   * @param data The session data to sign.
   * @returns The signed session data.
   */
  sign(data: Session): string {
    return sign(data, this.#privateKey, { algorithm: "RS256" });
  }

  /**
   * Verifies the signature of the session data.
   *
   * @param data The session data to verify.
   * @returns The verified session data.
   */
  verify(data: string): Session {
    return verify(data, this.#publicKey, { algorithms: ["RS256"], complete: false }) as Session;
  }
}
