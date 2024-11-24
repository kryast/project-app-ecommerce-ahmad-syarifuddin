--
-- PostgreSQL database dump
--

-- Dumped from database version 17.0 (Ubuntu 17.0-1.pgdg24.04+1)
-- Dumped by pg_dump version 17.0 (Ubuntu 17.0-1.pgdg24.04+1)

-- Started on 2024-11-24 15:37:35 WIB

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 244 (class 1259 OID 196609)
-- Name: address; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.address (
    id integer NOT NULL,
    auth_id integer,
    address text,
    flag_default boolean
);


ALTER TABLE public.address OWNER TO postgres;

--
-- TOC entry 243 (class 1259 OID 196608)
-- Name: address_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.address_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.address_id_seq OWNER TO postgres;

--
-- TOC entry 3565 (class 0 OID 0)
-- Dependencies: 243
-- Name: address_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.address_id_seq OWNED BY public.address.id;


--
-- TOC entry 218 (class 1259 OID 163842)
-- Name: auths; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.auths (
    id integer NOT NULL,
    name character varying(100),
    email character varying(100),
    phone character varying(20),
    password character varying(100),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone,
    token character varying(256),
    CONSTRAINT email_or_phone CHECK ((((email IS NOT NULL) AND (phone IS NULL)) OR ((email IS NULL) AND (phone IS NOT NULL))))
);


ALTER TABLE public.auths OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 163841)
-- Name: auths_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.auths_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.auths_id_seq OWNER TO postgres;

--
-- TOC entry 3566 (class 0 OID 0)
-- Dependencies: 217
-- Name: auths_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.auths_id_seq OWNED BY public.auths.id;


--
-- TOC entry 236 (class 1259 OID 180242)
-- Name: banners; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.banners (
    id integer NOT NULL,
    photo_id integer,
    title character varying(255),
    subtitle character varying(255),
    path_page character varying(255)
);


ALTER TABLE public.banners OWNER TO postgres;

--
-- TOC entry 235 (class 1259 OID 180241)
-- Name: banners_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.banners_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.banners_id_seq OWNER TO postgres;

--
-- TOC entry 3567 (class 0 OID 0)
-- Dependencies: 235
-- Name: banners_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.banners_id_seq OWNED BY public.banners.id;


--
-- TOC entry 222 (class 1259 OID 172056)
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying(100)
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 172055)
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO postgres;

--
-- TOC entry 3568 (class 0 OID 0)
-- Dependencies: 221
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- TOC entry 228 (class 1259 OID 172081)
-- Name: photos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.photos (
    id integer NOT NULL,
    photo_url character varying(100)
);


ALTER TABLE public.photos OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 172080)
-- Name: photos_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.photos_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.photos_id_seq OWNER TO postgres;

--
-- TOC entry 3569 (class 0 OID 0)
-- Dependencies: 227
-- Name: photos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.photos_id_seq OWNED BY public.photos.id;


--
-- TOC entry 230 (class 1259 OID 172108)
-- Name: product_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_details (
    id integer NOT NULL,
    product_id integer,
    category_id integer,
    photo_id integer
);


ALTER TABLE public.product_details OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 172107)
-- Name: product_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.product_details_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.product_details_id_seq OWNER TO postgres;

--
-- TOC entry 3570 (class 0 OID 0)
-- Dependencies: 229
-- Name: product_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.product_details_id_seq OWNED BY public.product_details.id;


--
-- TOC entry 220 (class 1259 OID 172033)
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id integer NOT NULL,
    name character varying(100),
    detail text,
    price numeric(10,2),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.products OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 172032)
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.products_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.products_id_seq OWNER TO postgres;

--
-- TOC entry 3571 (class 0 OID 0)
-- Dependencies: 219
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- TOC entry 242 (class 1259 OID 188429)
-- Name: promo_weekly; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.promo_weekly (
    id integer NOT NULL,
    product_details_id integer,
    start_week date,
    end_week date
);


ALTER TABLE public.promo_weekly OWNER TO postgres;

--
-- TOC entry 241 (class 1259 OID 188428)
-- Name: promo_weekly_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.promo_weekly_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.promo_weekly_id_seq OWNER TO postgres;

--
-- TOC entry 3572 (class 0 OID 0)
-- Dependencies: 241
-- Name: promo_weekly_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.promo_weekly_id_seq OWNED BY public.promo_weekly.id;


--
-- TOC entry 226 (class 1259 OID 172073)
-- Name: ratings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ratings (
    id integer NOT NULL,
    rating integer,
    CONSTRAINT ratings_rating_check CHECK (((rating >= 1) AND (rating <= 5)))
);


ALTER TABLE public.ratings OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 172072)
-- Name: ratings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ratings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ratings_id_seq OWNER TO postgres;

--
-- TOC entry 3573 (class 0 OID 0)
-- Dependencies: 225
-- Name: ratings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ratings_id_seq OWNED BY public.ratings.id;


--
-- TOC entry 240 (class 1259 OID 188417)
-- Name: recomments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.recomments (
    id integer NOT NULL,
    product_details_id integer
);


ALTER TABLE public.recomments OWNER TO postgres;

--
-- TOC entry 239 (class 1259 OID 188416)
-- Name: recomments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.recomments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.recomments_id_seq OWNER TO postgres;

--
-- TOC entry 3574 (class 0 OID 0)
-- Dependencies: 239
-- Name: recomments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.recomments_id_seq OWNED BY public.recomments.id;


--
-- TOC entry 234 (class 1259 OID 180225)
-- Name: reviews; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reviews (
    id integer NOT NULL,
    transaction_items_id integer,
    ratings_id integer
);


ALTER TABLE public.reviews OWNER TO postgres;

--
-- TOC entry 233 (class 1259 OID 180224)
-- Name: reviews_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.reviews_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.reviews_id_seq OWNER TO postgres;

--
-- TOC entry 3575 (class 0 OID 0)
-- Dependencies: 233
-- Name: reviews_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reviews_id_seq OWNED BY public.reviews.id;


--
-- TOC entry 232 (class 1259 OID 172130)
-- Name: transaction_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transaction_items (
    id integer NOT NULL,
    transaction_id integer,
    product_details_id integer,
    quantity integer DEFAULT 1,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone
);


ALTER TABLE public.transaction_items OWNER TO postgres;

--
-- TOC entry 231 (class 1259 OID 172129)
-- Name: transaction_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transaction_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.transaction_items_id_seq OWNER TO postgres;

--
-- TOC entry 3576 (class 0 OID 0)
-- Dependencies: 231
-- Name: transaction_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transaction_items_id_seq OWNED BY public.transaction_items.id;


--
-- TOC entry 224 (class 1259 OID 172063)
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    status character varying(20),
    CONSTRAINT transactions_status_check CHECK (((status)::text = ANY ((ARRAY['in_process'::character varying, 'completed'::character varying, 'canceled'::character varying])::text[])))
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 172062)
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.transactions_id_seq OWNER TO postgres;

--
-- TOC entry 3577 (class 0 OID 0)
-- Dependencies: 223
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- TOC entry 238 (class 1259 OID 180256)
-- Name: wishlists; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wishlists (
    id integer NOT NULL,
    product_details_id integer
);


ALTER TABLE public.wishlists OWNER TO postgres;

--
-- TOC entry 237 (class 1259 OID 180255)
-- Name: wishlists_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wishlists_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.wishlists_id_seq OWNER TO postgres;

--
-- TOC entry 3578 (class 0 OID 0)
-- Dependencies: 237
-- Name: wishlists_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wishlists_id_seq OWNED BY public.wishlists.id;


--
-- TOC entry 3339 (class 2604 OID 196612)
-- Name: address id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.address ALTER COLUMN id SET DEFAULT nextval('public.address_id_seq'::regclass);


--
-- TOC entry 3319 (class 2604 OID 163845)
-- Name: auths id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auths ALTER COLUMN id SET DEFAULT nextval('public.auths_id_seq'::regclass);


--
-- TOC entry 3335 (class 2604 OID 180245)
-- Name: banners id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banners ALTER COLUMN id SET DEFAULT nextval('public.banners_id_seq'::regclass);


--
-- TOC entry 3325 (class 2604 OID 172059)
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- TOC entry 3328 (class 2604 OID 172084)
-- Name: photos id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.photos ALTER COLUMN id SET DEFAULT nextval('public.photos_id_seq'::regclass);


--
-- TOC entry 3329 (class 2604 OID 172111)
-- Name: product_details id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_details ALTER COLUMN id SET DEFAULT nextval('public.product_details_id_seq'::regclass);


--
-- TOC entry 3322 (class 2604 OID 172036)
-- Name: products id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- TOC entry 3338 (class 2604 OID 188432)
-- Name: promo_weekly id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promo_weekly ALTER COLUMN id SET DEFAULT nextval('public.promo_weekly_id_seq'::regclass);


--
-- TOC entry 3327 (class 2604 OID 172076)
-- Name: ratings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ratings ALTER COLUMN id SET DEFAULT nextval('public.ratings_id_seq'::regclass);


--
-- TOC entry 3337 (class 2604 OID 188420)
-- Name: recomments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recomments ALTER COLUMN id SET DEFAULT nextval('public.recomments_id_seq'::regclass);


--
-- TOC entry 3334 (class 2604 OID 180228)
-- Name: reviews id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews ALTER COLUMN id SET DEFAULT nextval('public.reviews_id_seq'::regclass);


--
-- TOC entry 3330 (class 2604 OID 172133)
-- Name: transaction_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_items ALTER COLUMN id SET DEFAULT nextval('public.transaction_items_id_seq'::regclass);


--
-- TOC entry 3326 (class 2604 OID 172066)
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- TOC entry 3336 (class 2604 OID 180259)
-- Name: wishlists id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wishlists ALTER COLUMN id SET DEFAULT nextval('public.wishlists_id_seq'::regclass);


--
-- TOC entry 3559 (class 0 OID 196609)
-- Dependencies: 244
-- Data for Name: address; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.address (id, auth_id, address, flag_default) FROM stdin;
7	1	jalan baru no.1	t
2	45	jalan raya no.2	t
3	10	Alamat Ahmad 2	t
4	39	Alamat Ahmad 4	t
5	41	Alamat Ahmad 5	t
6	42	Alamat Ahmad 6	t
\.


--
-- TOC entry 3533 (class 0 OID 163842)
-- Dependencies: 218
-- Data for Name: auths; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.auths (id, name, email, phone, password, created_at, updated_at, deleted_at, token) FROM stdin;
39	Ahmad 4	ahmad4@example.com	\N	pass	2024-11-18 22:03:24.864895	2024-11-18 22:03:24.864895	\N	\N
41	Ahmad 5	ahmad5@example.com	\N	pass	2024-11-18 22:07:54.857111	2024-11-18 22:07:54.857111	\N	\N
42	Ahmad 6	ahmad6@example.com	\N	pass	2024-11-18 22:09:39.635326	2024-11-18 22:09:39.635326	\N	\N
43	Ahmad 7	ahmad7@example.com	\N	pass	2024-11-18 22:11:42.061365	2024-11-18 22:11:42.061365	\N	\N
31	ahmad 3	\N	123456789	pass	2024-11-18 21:55:56.134263	2024-11-18 21:55:56.134263	\N	31e6e0b5-f06f-4b69-ba97-d79596adf0d2
1	update ahmad	update2@example.com	\N	pass	2024-11-18 21:07:50.224564	2024-11-18 21:07:50.224564	\N	5feaadf0-9345-46f3-9f42-3b35e2e7ca50
10	Ahmad 2	ahmad2@example.com	\N	pass	2024-11-18 21:33:47.431629	2024-11-18 21:33:47.431629	\N	d8e042f5-90db-4f0f-bab7-d592a8180f66
45	Ahmad 8	ahmad8@example.com	\N	pass	2024-11-18 22:12:48.777336	2024-11-18 22:12:48.777336	\N	1020fced-4eb7-498f-8836-8c6bd46919e9
\.


--
-- TOC entry 3551 (class 0 OID 180242)
-- Dependencies: 236
-- Data for Name: banners; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.banners (id, photo_id, title, subtitle, path_page) FROM stdin;
1	1	Best Seller	See our best seller products	http://localhost:8080/best-seller
\.


--
-- TOC entry 3537 (class 0 OID 172056)
-- Dependencies: 222
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, name) FROM stdin;
1	Electronics
2	Fashion
3	Home and Kitchen
4	Beauty and Health
5	Sports and Outdoors
6	other
\.


--
-- TOC entry 3543 (class 0 OID 172081)
-- Dependencies: 228
-- Data for Name: photos; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.photos (id, photo_url) FROM stdin;
1	onepiece.jpg
\.


--
-- TOC entry 3545 (class 0 OID 172108)
-- Dependencies: 230
-- Data for Name: product_details; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product_details (id, product_id, category_id, photo_id) FROM stdin;
1	1	1	1
2	2	1	1
3	3	2	1
4	4	2	1
5	5	3	1
6	6	3	1
7	7	4	1
8	8	4	1
9	9	5	1
10	10	5	1
\.


--
-- TOC entry 3535 (class 0 OID 172033)
-- Dependencies: 220
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, name, detail, price, created_at, updated_at, deleted_at) FROM stdin;
1	Product 1	Description for product 1	100.00	2024-11-19 17:39:37.855245	2024-11-19 17:39:37.855245	\N
2	Product 2	Description for product 2	150.50	2024-11-19 17:39:37.855245	2024-11-19 17:39:37.855245	\N
3	Product 3	Description for product 3	200.00	2024-11-19 17:39:37.855245	2024-11-19 17:39:37.855245	\N
4	Product 4	Description for product 4	250.75	2024-11-19 17:39:37.855245	2024-11-19 17:39:37.855245	\N
5	Product 5	Description for product 5	300.00	2024-11-19 17:39:37.855245	2024-11-19 17:39:37.855245	\N
6	Product 6	Description for product 6	350.25	2024-11-19 17:39:37.855245	2024-11-19 17:39:37.855245	\N
7	Product 7	Description for product 7	400.00	2024-11-19 17:39:37.855245	2024-11-19 17:39:37.855245	\N
8	Product 8	Description for product 8	450.50	2024-11-19 17:39:37.855245	2024-11-19 17:39:37.855245	\N
9	Product 9	Description for product 9	500.00	2024-11-19 17:39:37.855245	2024-11-19 17:39:37.855245	\N
10	Product 10	Description for product 10	550.75	2024-11-19 17:39:37.855245	2024-11-19 17:39:37.855245	\N
\.


--
-- TOC entry 3557 (class 0 OID 188429)
-- Dependencies: 242
-- Data for Name: promo_weekly; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.promo_weekly (id, product_details_id, start_week, end_week) FROM stdin;
1	1	2024-11-21	2024-11-21
2	2	2024-11-21	2024-11-21
3	3	2024-11-21	2024-11-21
4	4	2024-12-01	2024-12-08
5	5	2024-12-01	2024-12-08
6	1	2024-11-22	2024-11-30
\.


--
-- TOC entry 3541 (class 0 OID 172073)
-- Dependencies: 226
-- Data for Name: ratings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ratings (id, rating) FROM stdin;
1	1
2	2
3	3
4	4
5	5
\.


--
-- TOC entry 3555 (class 0 OID 188417)
-- Dependencies: 240
-- Data for Name: recomments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.recomments (id, product_details_id) FROM stdin;
1	1
2	2
3	3
\.


--
-- TOC entry 3549 (class 0 OID 180225)
-- Dependencies: 234
-- Data for Name: reviews; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.reviews (id, transaction_items_id, ratings_id) FROM stdin;
1	4	5
2	5	4
3	6	3
4	12	5
5	13	5
6	14	5
\.


--
-- TOC entry 3547 (class 0 OID 172130)
-- Dependencies: 232
-- Data for Name: transaction_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transaction_items (id, transaction_id, product_details_id, quantity, created_at, updated_at, deleted_at) FROM stdin;
4	2	4	1	2024-11-19 20:28:36.476821	2024-11-19 20:28:36.476821	\N
5	2	5	1	2024-11-19 20:28:36.476821	2024-11-19 20:28:36.476821	\N
6	2	6	1	2024-11-19 20:28:36.476821	2024-11-19 20:28:36.476821	\N
7	3	7	1	2024-11-19 20:28:36.476821	2024-11-19 20:28:36.476821	\N
8	3	8	1	2024-11-19 20:28:36.476821	2024-11-19 20:28:36.476821	\N
9	3	9	1	2024-11-19 20:28:36.476821	2024-11-19 20:28:36.476821	\N
12	2	1	5	2024-11-20 00:26:56.286447	2024-11-20 00:26:56.286447	\N
13	2	2	3	2024-11-20 00:26:56.286447	2024-11-20 00:26:56.286447	\N
14	2	3	4	2024-11-20 00:26:56.286447	2024-11-20 00:26:56.286447	\N
15	2	1	2	2024-11-20 00:26:56.286447	2024-11-20 00:26:56.286447	\N
23	1	1	2	2024-11-23 04:41:42.332292	2024-11-23 04:41:42.332292	\N
22	1	2	4	2024-11-23 04:41:26.255369	2024-11-23 04:41:26.255369	\N
\.


--
-- TOC entry 3539 (class 0 OID 172063)
-- Dependencies: 224
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transactions (id, status) FROM stdin;
1	in_process
2	completed
3	canceled
\.


--
-- TOC entry 3553 (class 0 OID 180256)
-- Dependencies: 238
-- Data for Name: wishlists; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wishlists (id, product_details_id) FROM stdin;
3	1
6	2
7	2
\.


--
-- TOC entry 3579 (class 0 OID 0)
-- Dependencies: 243
-- Name: address_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.address_id_seq', 7, true);


--
-- TOC entry 3580 (class 0 OID 0)
-- Dependencies: 217
-- Name: auths_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.auths_id_seq', 45, true);


--
-- TOC entry 3581 (class 0 OID 0)
-- Dependencies: 235
-- Name: banners_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.banners_id_seq', 1, true);


--
-- TOC entry 3582 (class 0 OID 0)
-- Dependencies: 221
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 6, true);


--
-- TOC entry 3583 (class 0 OID 0)
-- Dependencies: 227
-- Name: photos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.photos_id_seq', 1, true);


--
-- TOC entry 3584 (class 0 OID 0)
-- Dependencies: 229
-- Name: product_details_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_details_id_seq', 10, true);


--
-- TOC entry 3585 (class 0 OID 0)
-- Dependencies: 219
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.products_id_seq', 10, true);


--
-- TOC entry 3586 (class 0 OID 0)
-- Dependencies: 241
-- Name: promo_weekly_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.promo_weekly_id_seq', 6, true);


--
-- TOC entry 3587 (class 0 OID 0)
-- Dependencies: 225
-- Name: ratings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ratings_id_seq', 5, true);


--
-- TOC entry 3588 (class 0 OID 0)
-- Dependencies: 239
-- Name: recomments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.recomments_id_seq', 3, true);


--
-- TOC entry 3589 (class 0 OID 0)
-- Dependencies: 233
-- Name: reviews_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.reviews_id_seq', 6, true);


--
-- TOC entry 3590 (class 0 OID 0)
-- Dependencies: 231
-- Name: transaction_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transaction_items_id_seq', 23, true);


--
-- TOC entry 3591 (class 0 OID 0)
-- Dependencies: 223
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 3, true);


--
-- TOC entry 3592 (class 0 OID 0)
-- Dependencies: 237
-- Name: wishlists_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wishlists_id_seq', 7, true);


--
-- TOC entry 3374 (class 2606 OID 196616)
-- Name: address address_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.address
    ADD CONSTRAINT address_pkey PRIMARY KEY (id);


--
-- TOC entry 3344 (class 2606 OID 163852)
-- Name: auths auths_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auths
    ADD CONSTRAINT auths_email_key UNIQUE (email);


--
-- TOC entry 3346 (class 2606 OID 163854)
-- Name: auths auths_phone_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auths
    ADD CONSTRAINT auths_phone_key UNIQUE (phone);


--
-- TOC entry 3348 (class 2606 OID 163850)
-- Name: auths auths_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auths
    ADD CONSTRAINT auths_pkey PRIMARY KEY (id);


--
-- TOC entry 3366 (class 2606 OID 180249)
-- Name: banners banners_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banners
    ADD CONSTRAINT banners_pkey PRIMARY KEY (id);


--
-- TOC entry 3352 (class 2606 OID 172061)
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- TOC entry 3358 (class 2606 OID 172086)
-- Name: photos photos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.photos
    ADD CONSTRAINT photos_pkey PRIMARY KEY (id);


--
-- TOC entry 3360 (class 2606 OID 172113)
-- Name: product_details product_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_details
    ADD CONSTRAINT product_details_pkey PRIMARY KEY (id);


--
-- TOC entry 3350 (class 2606 OID 172042)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- TOC entry 3372 (class 2606 OID 188434)
-- Name: promo_weekly promo_weekly_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promo_weekly
    ADD CONSTRAINT promo_weekly_pkey PRIMARY KEY (id);


--
-- TOC entry 3356 (class 2606 OID 172079)
-- Name: ratings ratings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ratings
    ADD CONSTRAINT ratings_pkey PRIMARY KEY (id);


--
-- TOC entry 3370 (class 2606 OID 188422)
-- Name: recomments recomments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recomments
    ADD CONSTRAINT recomments_pkey PRIMARY KEY (id);


--
-- TOC entry 3364 (class 2606 OID 180230)
-- Name: reviews reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (id);


--
-- TOC entry 3362 (class 2606 OID 172138)
-- Name: transaction_items transaction_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_items
    ADD CONSTRAINT transaction_items_pkey PRIMARY KEY (id);


--
-- TOC entry 3354 (class 2606 OID 172071)
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- TOC entry 3368 (class 2606 OID 180261)
-- Name: wishlists wishlists_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wishlists
    ADD CONSTRAINT wishlists_pkey PRIMARY KEY (id);


--
-- TOC entry 3386 (class 2606 OID 196617)
-- Name: address address_auth_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.address
    ADD CONSTRAINT address_auth_id_fkey FOREIGN KEY (auth_id) REFERENCES public.auths(id);


--
-- TOC entry 3382 (class 2606 OID 180250)
-- Name: banners banners_photo_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banners
    ADD CONSTRAINT banners_photo_id_fkey FOREIGN KEY (photo_id) REFERENCES public.photos(id);


--
-- TOC entry 3375 (class 2606 OID 172119)
-- Name: product_details product_details_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_details
    ADD CONSTRAINT product_details_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE CASCADE;


--
-- TOC entry 3376 (class 2606 OID 172124)
-- Name: product_details product_details_photo_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_details
    ADD CONSTRAINT product_details_photo_id_fkey FOREIGN KEY (photo_id) REFERENCES public.photos(id) ON DELETE CASCADE;


--
-- TOC entry 3377 (class 2606 OID 172114)
-- Name: product_details product_details_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_details
    ADD CONSTRAINT product_details_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE;


--
-- TOC entry 3385 (class 2606 OID 188435)
-- Name: promo_weekly promo_weekly_product_details_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.promo_weekly
    ADD CONSTRAINT promo_weekly_product_details_id_fkey FOREIGN KEY (product_details_id) REFERENCES public.product_details(id);


--
-- TOC entry 3384 (class 2606 OID 188423)
-- Name: recomments recomments_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.recomments
    ADD CONSTRAINT recomments_product_id_fkey FOREIGN KEY (product_details_id) REFERENCES public.product_details(id);


--
-- TOC entry 3380 (class 2606 OID 180236)
-- Name: reviews reviews_ratings_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_ratings_id_fkey FOREIGN KEY (ratings_id) REFERENCES public.ratings(id);


--
-- TOC entry 3381 (class 2606 OID 180231)
-- Name: reviews reviews_transaction_items_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_transaction_items_id_fkey FOREIGN KEY (transaction_items_id) REFERENCES public.transaction_items(id);


--
-- TOC entry 3378 (class 2606 OID 172144)
-- Name: transaction_items transaction_items_product_details_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_items
    ADD CONSTRAINT transaction_items_product_details_id_fkey FOREIGN KEY (product_details_id) REFERENCES public.product_details(id) ON DELETE CASCADE;


--
-- TOC entry 3379 (class 2606 OID 172139)
-- Name: transaction_items transaction_items_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_items
    ADD CONSTRAINT transaction_items_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES public.transactions(id) ON DELETE CASCADE;


--
-- TOC entry 3383 (class 2606 OID 180262)
-- Name: wishlists wishlists_product_details_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wishlists
    ADD CONSTRAINT wishlists_product_details_id_fkey FOREIGN KEY (product_details_id) REFERENCES public.product_details(id);


-- Completed on 2024-11-24 15:37:38 WIB

--
-- PostgreSQL database dump complete
--

