

export class Message {
    constructor(ID, Type, UserID, Username, Created_At, Body) {
        this.ID = ID;
        this.Type = Type;
        this.UserID = UserID;
        this.Username = Username;
        this.Created_At = Created_At
        this.Body = Body
    }
};

export class MessageRequest {
    constructor(Username, Body) {
        this.Username = Username;
        this.Body = Body;
    }
}