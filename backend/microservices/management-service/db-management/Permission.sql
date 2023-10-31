CREATE TABLE [dbo].[Permission] (
    [Id]        UNIQUEIDENTIFIER CONSTRAINT [DEFAULT_Permission_Id] DEFAULT (newid()) NOT NULL PRIMARY KEY CLUSTERED ([Id] ASC),
    [UserId]    UNIQUEIDENTIFIER NOT NULL,
    [RoleId]    UNIQUEIDENTIFIER NOT NULL,
    [CompanyId] UNIQUEIDENTIFIER NULL,
    CONSTRAINT [FK_Permission_Company] FOREIGN KEY ([CompanyId]) REFERENCES [dbo].[Company] ([Id]) ON DELETE CASCADE,
    CONSTRAINT [FK_Permission_Role] FOREIGN KEY ([RoleId]) REFERENCES [dbo].[Role] ([Id]),
    CONSTRAINT [FK_Permission_User] FOREIGN KEY ([UserId]) REFERENCES [dbo].[User] ([Id]) ON DELETE CASCADE
);
