CREATE TABLE [dbo].[User] (
    [Id]           UNIQUEIDENTIFIER CONSTRAINT [DEFAULT_User_Id] DEFAULT (newid()) NOT NULL PRIMARY KEY CLUSTERED ([Id] ASC),
    [Email]        NVARCHAR (255)   NOT NULL,
    [Username]     NVARCHAR (255)   NOT NULL,
    [PasswordHash] NVARCHAR (255)   NOT NULL,
    [Salt]         NVARCHAR (255)   NOT NULL,
    [RefreshToken] NVARCHAR (255)   NULL
);
