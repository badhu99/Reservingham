CREATE TABLE [dbo].[User] (
    [Id]             UNIQUEIDENTIFIER CONSTRAINT [DEFAULT_User_Id] DEFAULT (newid()) NOT NULL PRIMARY KEY CLUSTERED ([Id] ASC),
    [Email]          NVARCHAR (255)   NOT NULL,
    [Username]       NVARCHAR (255)   NULL,
    [PasswordHash]   NVARCHAR (255)   NULL,
    [Salt]           NVARCHAR (255)   NULL,
    [RefreshToken]   NVARCHAR (255)   NULL,
    [Firstname]      NVARCHAR (255)   NULL,
    [Lastname]       NVARCHAR (255)   NULL,
    [Code]           NVARCHAR (20)    NULL,
    [PasswordChange] BIT              NOT NULL
);
