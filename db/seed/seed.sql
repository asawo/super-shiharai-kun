--
-- PostgreSQL database dump
--

-- Dumped from database version 12.0
-- Dumped by pg_dump version 15.6 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: invoices-user
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO "invoices-user";

--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: invoices-user
--

COMMENT ON SCHEMA public IS '';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: bank_accounts; Type: TABLE; Schema: public; Owner: invoices-user
--

CREATE TABLE public.bank_accounts (
    id bigint NOT NULL,
    service_provider_id bigint NOT NULL,
    bank_name text NOT NULL,
    branch_name text NOT NULL,
    account_number text NOT NULL,
    account_name text NOT NULL
);


ALTER TABLE public.bank_accounts OWNER TO "invoices-user";

--
-- Name: bank_accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: invoices-user
--

CREATE SEQUENCE public.bank_accounts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bank_accounts_id_seq OWNER TO "invoices-user";

--
-- Name: bank_accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: invoices-user
--

ALTER SEQUENCE public.bank_accounts_id_seq OWNED BY public.bank_accounts.id;


--
-- Name: companies; Type: TABLE; Schema: public; Owner: invoices-user
--

CREATE TABLE public.companies (
    id bigint NOT NULL,
    name text NOT NULL,
    representative text NOT NULL,
    phone_number text NOT NULL,
    postal_code text NOT NULL,
    address text NOT NULL
);


ALTER TABLE public.companies OWNER TO "invoices-user";

--
-- Name: companies_id_seq; Type: SEQUENCE; Schema: public; Owner: invoices-user
--

CREATE SEQUENCE public.companies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.companies_id_seq OWNER TO "invoices-user";

--
-- Name: companies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: invoices-user
--

ALTER SEQUENCE public.companies_id_seq OWNED BY public.companies.id;


--
-- Name: invoices; Type: TABLE; Schema: public; Owner: invoices-user
--

CREATE TABLE public.invoices (
    id text NOT NULL,
    issue_date timestamp with time zone,
    payment_amount numeric,
    commission numeric,
    commission_rate numeric,
    tax numeric,
    tax_rate numeric,
    amount numeric,
    due_date timestamp with time zone,
    company_id bigint,
    service_provider_id bigint,
    status text
);


ALTER TABLE public.invoices OWNER TO "invoices-user";

--
-- Name: service_providers; Type: TABLE; Schema: public; Owner: invoices-user
--

CREATE TABLE public.service_providers (
    id bigint NOT NULL,
    company_id bigint NOT NULL,
    name text NOT NULL,
    representative text NOT NULL,
    phone_number text NOT NULL,
    postal_code text NOT NULL,
    address text NOT NULL
);


ALTER TABLE public.service_providers OWNER TO "invoices-user";

--
-- Name: service_providers_id_seq; Type: SEQUENCE; Schema: public; Owner: invoices-user
--

CREATE SEQUENCE public.service_providers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.service_providers_id_seq OWNER TO "invoices-user";

--
-- Name: service_providers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: invoices-user
--

ALTER SEQUENCE public.service_providers_id_seq OWNED BY public.service_providers.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: invoices-user
--

CREATE TABLE public.users (
    company_id bigint NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    password text NOT NULL
);


ALTER TABLE public.users OWNER TO "invoices-user";

--
-- Name: bank_accounts id; Type: DEFAULT; Schema: public; Owner: invoices-user
--

ALTER TABLE ONLY public.bank_accounts ALTER COLUMN id SET DEFAULT nextval('public.bank_accounts_id_seq'::regclass);


--
-- Name: companies id; Type: DEFAULT; Schema: public; Owner: invoices-user
--

ALTER TABLE ONLY public.companies ALTER COLUMN id SET DEFAULT nextval('public.companies_id_seq'::regclass);


--
-- Name: service_providers id; Type: DEFAULT; Schema: public; Owner: invoices-user
--

ALTER TABLE ONLY public.service_providers ALTER COLUMN id SET DEFAULT nextval('public.service_providers_id_seq'::regclass);


--
-- Data for Name: bank_accounts; Type: TABLE DATA; Schema: public; Owner: invoices-user
--

COPY public.bank_accounts (id, service_provider_id, bank_name, branch_name, account_number, account_name) FROM stdin;
1	1	Bank of Example	Main Branch	1234567890	Jane Smith
2	2	Bank of Example 2	Main Branch	1234567890	Jane Smith
3	3	Bank of Example	Main Branch	1234567890	Jane Smith
\.


--
-- Data for Name: companies; Type: TABLE DATA; Schema: public; Owner: invoices-user
--

COPY public.companies (id, name, representative, phone_number, postal_code, address) FROM stdin;
1	Example Corporation	John Doe	123456789	12345	123 Main Street
2	Example Corporation 3	Ben Doe	123456789	12345	123 Main Street
\.


--
-- Data for Name: invoices; Type: TABLE DATA; Schema: public; Owner: invoices-user
--

COPY public.invoices (id, issue_date, payment_amount, commission, commission_rate, tax, tax_rate, amount, due_date, company_id, service_provider_id, status) FROM stdin;
b796e13d-bd76-46a6-af4e-159f8e19587f	2024-02-15 02:12:57.06461+00	100	4	0.04	1.1	0.1	104.4	2024-03-16 02:12:57.06461+00	1	1	OUTSTANDING
c5358cbc-aef7-4398-8b36-6f77f40cfa55	2024-02-15 02:12:57.067696+00	100	4	0.04	1.1	0.1	104.4	2024-03-16 02:12:57.067696+00	2	2	OUTSTANDING
\.


--
-- Data for Name: service_providers; Type: TABLE DATA; Schema: public; Owner: invoices-user
--

COPY public.service_providers (id, company_id, name, representative, phone_number, postal_code, address) FROM stdin;
1	1	serviceProvider Corporation	Jane Smith	987654321	54321	456 Oak Avenue
2	2	serviceProvider Corporation 2	Jane Smith	987654321	54321	456 Oak Avenue
3	2	serviceProvider Corporation 3	Jane Smith	987654321	54321	456 Oak Avenue
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: invoices-user
--

COPY public.users (company_id, name, email, password) FROM stdin;
1	Alice	alice@example.com	$2a$10$AwlfXUFKrIhy3O/Sxo0OHeEFcV1wBf9A32/nVD4rzz.yPuEKfeayu
2	Arthur	arthur@example.com	$2a$10$8KWLi1dmZY3xSCfdBCaoOeThzqmrf35ksNdbQUYSsfeUtFIQcOI4W
\.


--
-- Name: bank_accounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: invoices-user
--

SELECT pg_catalog.setval('public.bank_accounts_id_seq', 3, true);


--
-- Name: companies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: invoices-user
--

SELECT pg_catalog.setval('public.companies_id_seq', 1, false);


--
-- Name: service_providers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: invoices-user
--

SELECT pg_catalog.setval('public.service_providers_id_seq', 3, true);


--
-- Name: bank_accounts bank_accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: invoices-user
--

ALTER TABLE ONLY public.bank_accounts
    ADD CONSTRAINT bank_accounts_pkey PRIMARY KEY (id);


--
-- Name: companies companies_pkey; Type: CONSTRAINT; Schema: public; Owner: invoices-user
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT companies_pkey PRIMARY KEY (id);


--
-- Name: invoices invoices_pkey; Type: CONSTRAINT; Schema: public; Owner: invoices-user
--

ALTER TABLE ONLY public.invoices
    ADD CONSTRAINT invoices_pkey PRIMARY KEY (id);


--
-- Name: service_providers service_providers_pkey; Type: CONSTRAINT; Schema: public; Owner: invoices-user
--

ALTER TABLE ONLY public.service_providers
    ADD CONSTRAINT service_providers_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: invoices-user
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: invoices fk_companies_invoices; Type: FK CONSTRAINT; Schema: public; Owner: invoices-user
--

ALTER TABLE ONLY public.invoices
    ADD CONSTRAINT fk_companies_invoices FOREIGN KEY (company_id) REFERENCES public.companies(id);


--
-- Name: service_providers fk_companies_service_providers; Type: FK CONSTRAINT; Schema: public; Owner: invoices-user
--

ALTER TABLE ONLY public.service_providers
    ADD CONSTRAINT fk_companies_service_providers FOREIGN KEY (company_id) REFERENCES public.companies(id);


--
-- Name: users fk_companies_users; Type: FK CONSTRAINT; Schema: public; Owner: invoices-user
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_companies_users FOREIGN KEY (company_id) REFERENCES public.companies(id);


--
-- Name: bank_accounts fk_service_providers_bank_accounts; Type: FK CONSTRAINT; Schema: public; Owner: invoices-user
--

ALTER TABLE ONLY public.bank_accounts
    ADD CONSTRAINT fk_service_providers_bank_accounts FOREIGN KEY (service_provider_id) REFERENCES public.service_providers(id);


--
-- Name: invoices fk_service_providers_invoices; Type: FK CONSTRAINT; Schema: public; Owner: invoices-user
--

ALTER TABLE ONLY public.invoices
    ADD CONSTRAINT fk_service_providers_invoices FOREIGN KEY (service_provider_id) REFERENCES public.service_providers(id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: invoices-user
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;


--
-- PostgreSQL database dump complete
--

