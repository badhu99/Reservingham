CREATE TABLE [dbo].[Role] (
    [Id]   UNIQUEIDENTIFIER CONSTRAINT [DEFAULT_Role_Id] DEFAULT (newid()) NOT NULL PRIMARY KEY CLUSTERED ([Id] ASC),
    [Name] NVARCHAR (255)   NOT NULL
);
