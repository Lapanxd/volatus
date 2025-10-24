import {User} from "./user.ts";

export interface PendingHandshakesOutput {
    session_id: string;
    from_user_id: number;
    sdp_offer: string;
}

export interface PendingHandshakeWithUser {
    sessionId: string;
    fromUser: User;
    sdpOffer: string;
}