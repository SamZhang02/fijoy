// @generated
// This file is automatically generated by Kanel. Do not modify manually.

import { fijoyTransactionType, type default as FijoyTransactionType } from './FijoyTransactionType';
import { fijoyAccountId, type FijoyAccountId } from './FijoyAccount';
import { fijoyUserId, type FijoyUserId } from './FijoyUser';
import { fijoyWorkspaceId, type FijoyWorkspaceId } from './FijoyWorkspace';
import { fijoyCategoryId, type FijoyCategoryId } from './FijoyCategory';
import { z } from 'zod';

export type FijoyTransactionId = string;

/** Represents the table public.fijoy_transaction */
export default interface FijoyTransaction {
  /** Database type: pg_catalog.text */
  id: FijoyTransactionId;

  /** Database type: public.fijoy_transaction_type */
  transactionType: FijoyTransactionType;

  /** Database type: pg_catalog.numeric */
  amount: string;

  /** Database type: pg_catalog.text */
  currency: string;

  /** Database type: pg_catalog.text */
  accountId: FijoyAccountId | null;

  /** Database type: pg_catalog.text */
  userId: FijoyUserId;

  /** Database type: pg_catalog.text */
  workspaceId: FijoyWorkspaceId;

  /** Database type: pg_catalog.timestamptz */
  datetime: Date;

  /** Database type: pg_catalog.text */
  note: string | null;

  /** Database type: pg_catalog.text */
  categoryId: FijoyCategoryId | null;

  /** Database type: pg_catalog.text */
  entity: string | null;
}

/** Represents the initializer for the table public.fijoy_transaction */
export interface FijoyTransactionInitializer {
  /** Database type: pg_catalog.text */
  id: FijoyTransactionId;

  /** Database type: public.fijoy_transaction_type */
  transactionType: FijoyTransactionType;

  /** Database type: pg_catalog.numeric */
  amount: string;

  /** Database type: pg_catalog.text */
  currency: string;

  /** Database type: pg_catalog.text */
  accountId?: FijoyAccountId | null;

  /** Database type: pg_catalog.text */
  userId: FijoyUserId;

  /** Database type: pg_catalog.text */
  workspaceId: FijoyWorkspaceId;

  /** Database type: pg_catalog.timestamptz */
  datetime: Date;

  /** Database type: pg_catalog.text */
  note?: string | null;

  /** Database type: pg_catalog.text */
  categoryId?: FijoyCategoryId | null;

  /** Database type: pg_catalog.text */
  entity?: string | null;
}

/** Represents the mutator for the table public.fijoy_transaction */
export interface FijoyTransactionMutator {
  /** Database type: pg_catalog.text */
  id?: FijoyTransactionId;

  /** Database type: public.fijoy_transaction_type */
  transactionType?: FijoyTransactionType;

  /** Database type: pg_catalog.numeric */
  amount?: string;

  /** Database type: pg_catalog.text */
  currency?: string;

  /** Database type: pg_catalog.text */
  accountId?: FijoyAccountId | null;

  /** Database type: pg_catalog.text */
  userId?: FijoyUserId;

  /** Database type: pg_catalog.text */
  workspaceId?: FijoyWorkspaceId;

  /** Database type: pg_catalog.timestamptz */
  datetime?: Date;

  /** Database type: pg_catalog.text */
  note?: string | null;

  /** Database type: pg_catalog.text */
  categoryId?: FijoyCategoryId | null;

  /** Database type: pg_catalog.text */
  entity?: string | null;
}

export const fijoyTransactionId = z.string();

export const fijoyTransaction = z.object({
  id: fijoyTransactionId,
  transactionType: fijoyTransactionType,
  amount: z.string(),
  currency: z.string(),
  accountId: fijoyAccountId.nullable(),
  userId: fijoyUserId,
  workspaceId: fijoyWorkspaceId,
  datetime: z.coerce.date(),
  note: z.string().nullable(),
  categoryId: fijoyCategoryId.nullable(),
  entity: z.string().nullable(),
});

export const fijoyTransactionInitializer = z.object({
  id: fijoyTransactionId,
  transactionType: fijoyTransactionType,
  amount: z.string(),
  currency: z.string(),
  accountId: fijoyAccountId.optional().nullable(),
  userId: fijoyUserId,
  workspaceId: fijoyWorkspaceId,
  datetime: z.coerce.date(),
  note: z.string().optional().nullable(),
  categoryId: fijoyCategoryId.optional().nullable(),
  entity: z.string().optional().nullable(),
});

export const fijoyTransactionMutator = z.object({
  id: fijoyTransactionId.optional(),
  transactionType: fijoyTransactionType.optional(),
  amount: z.string().optional(),
  currency: z.string().optional(),
  accountId: fijoyAccountId.optional().nullable(),
  userId: fijoyUserId.optional(),
  workspaceId: fijoyWorkspaceId.optional(),
  datetime: z.coerce.date().optional(),
  note: z.string().optional().nullable(),
  categoryId: fijoyCategoryId.optional().nullable(),
  entity: z.string().optional().nullable(),
});
