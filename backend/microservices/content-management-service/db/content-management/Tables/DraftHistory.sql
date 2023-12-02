CREATE TABLE [dbo].[DraftHistory] (
    [Id]      UNIQUEIDENTIFIER CONSTRAINT [DEFAULT_DraftHistory_Id] DEFAULT (newid()) NOT NULL PRIMARY KEY CLUSTERED ([Id] ASC),
    [DraftId] UNIQUEIDENTIFIER NOT NULL,
    [FileId]  UNIQUEIDENTIFIER NOT NULL,
    [UserId]  UNIQUEIDENTIFIER NOT NULL,
    [Date]    DATETIME         NOT NULL,
    [Title]   NVARCHAR (255)   NOT NULL,
    [Message] NVARCHAR (255)   NULL,
    CONSTRAINT [FK_DraftHistory_Document] FOREIGN KEY ([FileId]) REFERENCES [dbo].[Document] ([Id]),
    CONSTRAINT [FK_DraftHistory_Draft] FOREIGN KEY ([DraftId]) REFERENCES [dbo].[Draft] ([Id])
);
