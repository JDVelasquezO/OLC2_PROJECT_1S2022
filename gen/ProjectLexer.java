// Generated from /home/jdvelasquezo/GolandProjects/OLC2_Project1/server/ProjectLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ProjectLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PRINTLN=1, PRINT=2, DECLARAR=3, MUT=4, RSTRING=5, RINTEGER=6, RREAL=7, 
		RBOOLEAN=8, RCHAR=9, RSTR=10, RAS=11, RIF=12, RELSE=13, RWHILE=14, POWI=15, 
		POWF=16, TOSTRING=17, TOOWNED=18, MULTILINE=19, INLINE=20, INTEGER=21, 
		FLOAT=22, CHAR=23, STRING=24, BOOLEAN=25, ID=26, EQUAL=27, DOT=28, SEMICOLON=29, 
		COMMA=30, COLON=31, ADMIRATION=32, REFERENCE=33, HERITAGE=34, LEFT_PAR=35, 
		RIGHT_PAR=36, LEFT_KEY=37, RIGHT_KEY=38, LEFT_BRACKET=39, RIGHT_BRACKET=40, 
		EQUEAL_EQUAL=41, NOTEQUAL=42, GREATER_THAN=43, LESS_THAN=44, GREATER_EQUALTHAN=45, 
		LESS_EQUEALTHAN=46, MUL=47, DIV=48, MOD=49, ADD=50, SUB=51, AND=52, OR=53, 
		WHITESPACE=54;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL", 
			"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "RWHILE", "POWI", 
			"POWF", "TOSTRING", "TOOWNED", "MULTILINE", "INLINE", "INTEGER", "FLOAT", 
			"CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT", "SEMICOLON", "COMMA", 
			"COLON", "ADMIRATION", "REFERENCE", "HERITAGE", "LEFT_PAR", "RIGHT_PAR", 
			"LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL", 
			"NOTEQUAL", "GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN", 
			"MUL", "DIV", "MOD", "ADD", "SUB", "AND", "OR", "WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'println'", "'print'", "'let'", "'mut'", "'String'", "'i64'", 
			"'f64'", "'bool'", "'char'", "'&str'", "'as'", "'if'", "'else'", "'while'", 
			"'pow'", "'powf'", "'to_string'", "'to_owned'", null, null, null, null, 
			null, null, null, null, "'='", "'.'", "';'", "','", "':'", "'!'", "'&'", 
			"'::'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'=='", "'!='", "'>'", 
			"'<'", "'>='", "'<='", "'*'", "'/'", "'%'", "'+'", "'-'", "'&&'", "'||'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL", 
			"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "RWHILE", "POWI", 
			"POWF", "TOSTRING", "TOOWNED", "MULTILINE", "INLINE", "INTEGER", "FLOAT", 
			"CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT", "SEMICOLON", "COMMA", 
			"COLON", "ADMIRATION", "REFERENCE", "HERITAGE", "LEFT_PAR", "RIGHT_PAR", 
			"LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL", 
			"NOTEQUAL", "GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN", 
			"MUL", "DIV", "MOD", "ADD", "SUB", "AND", "OR", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public ProjectLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ProjectLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\28\u015f\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3"+
		"\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3"+
		"\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\r\3"+
		"\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3"+
		"\20\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3"+
		"\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3"+
		"\24\6\24\u00d7\n\24\r\24\16\24\u00d8\3\24\3\24\3\24\3\24\3\24\3\25\3\25"+
		"\3\25\3\25\6\25\u00e4\n\25\r\25\16\25\u00e5\3\25\3\25\3\26\6\26\u00eb"+
		"\n\26\r\26\16\26\u00ec\3\27\6\27\u00f0\n\27\r\27\16\27\u00f1\3\27\3\27"+
		"\6\27\u00f6\n\27\r\27\16\27\u00f7\3\30\3\30\3\30\3\30\3\31\3\31\7\31\u0100"+
		"\n\31\f\31\16\31\u0103\13\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3"+
		"\32\3\32\3\32\5\32\u0110\n\32\3\33\3\33\7\33\u0114\n\33\f\33\16\33\u0117"+
		"\13\33\3\34\3\34\3\35\3\35\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#"+
		"\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)\3*\3*\3*\3+\3+\3+\3,\3,\3"+
		"-\3-\3.\3.\3.\3/\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3\64\3"+
		"\64\3\65\3\65\3\65\3\66\3\66\3\66\3\67\6\67\u0157\n\67\r\67\16\67\u0158"+
		"\3\67\3\67\38\38\38\2\29\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f"+
		"\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63"+
		"\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62"+
		"c\63e\64g\65i\66k\67m8o\2\3\2\n\3\2\61\61\3\2\f\f\3\2\62;\3\2$$\7\2C\\"+
		"aac|\u00d3\u00d3\u00f3\u00f3\7\2\62;C\\c|\u00d3\u00d3\u00f3\u00f3\5\2"+
		"\13\f\17\17\"\"\t\2\"#%%--/\60<<BB]_\2\u0166\2\3\3\2\2\2\2\5\3\2\2\2\2"+
		"\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2"+
		"\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2"+
		"\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2"+
		"\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2"+
		"\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2"+
		"\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2"+
		"M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3"+
		"\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2"+
		"\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\3q\3\2\2\2\5y\3\2\2\2\7"+
		"\177\3\2\2\2\t\u0083\3\2\2\2\13\u0087\3\2\2\2\r\u008e\3\2\2\2\17\u0092"+
		"\3\2\2\2\21\u0096\3\2\2\2\23\u009b\3\2\2\2\25\u00a0\3\2\2\2\27\u00a5\3"+
		"\2\2\2\31\u00a8\3\2\2\2\33\u00ab\3\2\2\2\35\u00b0\3\2\2\2\37\u00b6\3\2"+
		"\2\2!\u00ba\3\2\2\2#\u00bf\3\2\2\2%\u00c9\3\2\2\2\'\u00d2\3\2\2\2)\u00df"+
		"\3\2\2\2+\u00ea\3\2\2\2-\u00ef\3\2\2\2/\u00f9\3\2\2\2\61\u00fd\3\2\2\2"+
		"\63\u010f\3\2\2\2\65\u0111\3\2\2\2\67\u0118\3\2\2\29\u011a\3\2\2\2;\u011c"+
		"\3\2\2\2=\u011e\3\2\2\2?\u0120\3\2\2\2A\u0122\3\2\2\2C\u0124\3\2\2\2E"+
		"\u0126\3\2\2\2G\u0129\3\2\2\2I\u012b\3\2\2\2K\u012d\3\2\2\2M\u012f\3\2"+
		"\2\2O\u0131\3\2\2\2Q\u0133\3\2\2\2S\u0135\3\2\2\2U\u0138\3\2\2\2W\u013b"+
		"\3\2\2\2Y\u013d\3\2\2\2[\u013f\3\2\2\2]\u0142\3\2\2\2_\u0145\3\2\2\2a"+
		"\u0147\3\2\2\2c\u0149\3\2\2\2e\u014b\3\2\2\2g\u014d\3\2\2\2i\u014f\3\2"+
		"\2\2k\u0152\3\2\2\2m\u0156\3\2\2\2o\u015c\3\2\2\2qr\7r\2\2rs\7t\2\2st"+
		"\7k\2\2tu\7p\2\2uv\7v\2\2vw\7n\2\2wx\7p\2\2x\4\3\2\2\2yz\7r\2\2z{\7t\2"+
		"\2{|\7k\2\2|}\7p\2\2}~\7v\2\2~\6\3\2\2\2\177\u0080\7n\2\2\u0080\u0081"+
		"\7g\2\2\u0081\u0082\7v\2\2\u0082\b\3\2\2\2\u0083\u0084\7o\2\2\u0084\u0085"+
		"\7w\2\2\u0085\u0086\7v\2\2\u0086\n\3\2\2\2\u0087\u0088\7U\2\2\u0088\u0089"+
		"\7v\2\2\u0089\u008a\7t\2\2\u008a\u008b\7k\2\2\u008b\u008c\7p\2\2\u008c"+
		"\u008d\7i\2\2\u008d\f\3\2\2\2\u008e\u008f\7k\2\2\u008f\u0090\78\2\2\u0090"+
		"\u0091\7\66\2\2\u0091\16\3\2\2\2\u0092\u0093\7h\2\2\u0093\u0094\78\2\2"+
		"\u0094\u0095\7\66\2\2\u0095\20\3\2\2\2\u0096\u0097\7d\2\2\u0097\u0098"+
		"\7q\2\2\u0098\u0099\7q\2\2\u0099\u009a\7n\2\2\u009a\22\3\2\2\2\u009b\u009c"+
		"\7e\2\2\u009c\u009d\7j\2\2\u009d\u009e\7c\2\2\u009e\u009f\7t\2\2\u009f"+
		"\24\3\2\2\2\u00a0\u00a1\7(\2\2\u00a1\u00a2\7u\2\2\u00a2\u00a3\7v\2\2\u00a3"+
		"\u00a4\7t\2\2\u00a4\26\3\2\2\2\u00a5\u00a6\7c\2\2\u00a6\u00a7\7u\2\2\u00a7"+
		"\30\3\2\2\2\u00a8\u00a9\7k\2\2\u00a9\u00aa\7h\2\2\u00aa\32\3\2\2\2\u00ab"+
		"\u00ac\7g\2\2\u00ac\u00ad\7n\2\2\u00ad\u00ae\7u\2\2\u00ae\u00af\7g\2\2"+
		"\u00af\34\3\2\2\2\u00b0\u00b1\7y\2\2\u00b1\u00b2\7j\2\2\u00b2\u00b3\7"+
		"k\2\2\u00b3\u00b4\7n\2\2\u00b4\u00b5\7g\2\2\u00b5\36\3\2\2\2\u00b6\u00b7"+
		"\7r\2\2\u00b7\u00b8\7q\2\2\u00b8\u00b9\7y\2\2\u00b9 \3\2\2\2\u00ba\u00bb"+
		"\7r\2\2\u00bb\u00bc\7q\2\2\u00bc\u00bd\7y\2\2\u00bd\u00be\7h\2\2\u00be"+
		"\"\3\2\2\2\u00bf\u00c0\7v\2\2\u00c0\u00c1\7q\2\2\u00c1\u00c2\7a\2\2\u00c2"+
		"\u00c3\7u\2\2\u00c3\u00c4\7v\2\2\u00c4\u00c5\7t\2\2\u00c5\u00c6\7k\2\2"+
		"\u00c6\u00c7\7p\2\2\u00c7\u00c8\7i\2\2\u00c8$\3\2\2\2\u00c9\u00ca\7v\2"+
		"\2\u00ca\u00cb\7q\2\2\u00cb\u00cc\7a\2\2\u00cc\u00cd\7q\2\2\u00cd\u00ce"+
		"\7y\2\2\u00ce\u00cf\7p\2\2\u00cf\u00d0\7g\2\2\u00d0\u00d1\7f\2\2\u00d1"+
		"&\3\2\2\2\u00d2\u00d3\7\61\2\2\u00d3\u00d4\7,\2\2\u00d4\u00d6\3\2\2\2"+
		"\u00d5\u00d7\n\2\2\2\u00d6\u00d5\3\2\2\2\u00d7\u00d8\3\2\2\2\u00d8\u00d6"+
		"\3\2\2\2\u00d8\u00d9\3\2\2\2\u00d9\u00da\3\2\2\2\u00da\u00db\7,\2\2\u00db"+
		"\u00dc\7\61\2\2\u00dc\u00dd\3\2\2\2\u00dd\u00de\b\24\2\2\u00de(\3\2\2"+
		"\2\u00df\u00e0\7\61\2\2\u00e0\u00e1\7\61\2\2\u00e1\u00e3\3\2\2\2\u00e2"+
		"\u00e4\n\3\2\2\u00e3\u00e2\3\2\2\2\u00e4\u00e5\3\2\2\2\u00e5\u00e3\3\2"+
		"\2\2\u00e5\u00e6\3\2\2\2\u00e6\u00e7\3\2\2\2\u00e7\u00e8\b\25\2\2\u00e8"+
		"*\3\2\2\2\u00e9\u00eb\t\4\2\2\u00ea\u00e9\3\2\2\2\u00eb\u00ec\3\2\2\2"+
		"\u00ec\u00ea\3\2\2\2\u00ec\u00ed\3\2\2\2\u00ed,\3\2\2\2\u00ee\u00f0\t"+
		"\4\2\2\u00ef\u00ee\3\2\2\2\u00f0\u00f1\3\2\2\2\u00f1\u00ef\3\2\2\2\u00f1"+
		"\u00f2\3\2\2\2\u00f2\u00f3\3\2\2\2\u00f3\u00f5\7\60\2\2\u00f4\u00f6\t"+
		"\4\2\2\u00f5\u00f4\3\2\2\2\u00f6\u00f7\3\2\2\2\u00f7\u00f5\3\2\2\2\u00f7"+
		"\u00f8\3\2\2\2\u00f8.\3\2\2\2\u00f9\u00fa\7)\2\2\u00fa\u00fb\n\5\2\2\u00fb"+
		"\u00fc\7)\2\2\u00fc\60\3\2\2\2\u00fd\u0101\7$\2\2\u00fe\u0100\n\5\2\2"+
		"\u00ff\u00fe\3\2\2\2\u0100\u0103\3\2\2\2\u0101\u00ff\3\2\2\2\u0101\u0102"+
		"\3\2\2\2\u0102\u0104\3\2\2\2\u0103\u0101\3\2\2\2\u0104\u0105\7$\2\2\u0105"+
		"\62\3\2\2\2\u0106\u0107\7v\2\2\u0107\u0108\7t\2\2\u0108\u0109\7w\2\2\u0109"+
		"\u0110\7g\2\2\u010a\u010b\7h\2\2\u010b\u010c\7c\2\2\u010c\u010d\7n\2\2"+
		"\u010d\u010e\7u\2\2\u010e\u0110\7g\2\2\u010f\u0106\3\2\2\2\u010f\u010a"+
		"\3\2\2\2\u0110\64\3\2\2\2\u0111\u0115\t\6\2\2\u0112\u0114\t\7\2\2\u0113"+
		"\u0112\3\2\2\2\u0114\u0117\3\2\2\2\u0115\u0113\3\2\2\2\u0115\u0116\3\2"+
		"\2\2\u0116\66\3\2\2\2\u0117\u0115\3\2\2\2\u0118\u0119\7?\2\2\u01198\3"+
		"\2\2\2\u011a\u011b\7\60\2\2\u011b:\3\2\2\2\u011c\u011d\7=\2\2\u011d<\3"+
		"\2\2\2\u011e\u011f\7.\2\2\u011f>\3\2\2\2\u0120\u0121\7<\2\2\u0121@\3\2"+
		"\2\2\u0122\u0123\7#\2\2\u0123B\3\2\2\2\u0124\u0125\7(\2\2\u0125D\3\2\2"+
		"\2\u0126\u0127\7<\2\2\u0127\u0128\7<\2\2\u0128F\3\2\2\2\u0129\u012a\7"+
		"*\2\2\u012aH\3\2\2\2\u012b\u012c\7+\2\2\u012cJ\3\2\2\2\u012d\u012e\7}"+
		"\2\2\u012eL\3\2\2\2\u012f\u0130\7\177\2\2\u0130N\3\2\2\2\u0131\u0132\7"+
		"]\2\2\u0132P\3\2\2\2\u0133\u0134\7_\2\2\u0134R\3\2\2\2\u0135\u0136\7?"+
		"\2\2\u0136\u0137\7?\2\2\u0137T\3\2\2\2\u0138\u0139\7#\2\2\u0139\u013a"+
		"\7?\2\2\u013aV\3\2\2\2\u013b\u013c\7@\2\2\u013cX\3\2\2\2\u013d\u013e\7"+
		">\2\2\u013eZ\3\2\2\2\u013f\u0140\7@\2\2\u0140\u0141\7?\2\2\u0141\\\3\2"+
		"\2\2\u0142\u0143\7>\2\2\u0143\u0144\7?\2\2\u0144^\3\2\2\2\u0145\u0146"+
		"\7,\2\2\u0146`\3\2\2\2\u0147\u0148\7\61\2\2\u0148b\3\2\2\2\u0149\u014a"+
		"\7\'\2\2\u014ad\3\2\2\2\u014b\u014c\7-\2\2\u014cf\3\2\2\2\u014d\u014e"+
		"\7/\2\2\u014eh\3\2\2\2\u014f\u0150\7(\2\2\u0150\u0151\7(\2\2\u0151j\3"+
		"\2\2\2\u0152\u0153\7~\2\2\u0153\u0154\7~\2\2\u0154l\3\2\2\2\u0155\u0157"+
		"\t\b\2\2\u0156\u0155\3\2\2\2\u0157\u0158\3\2\2\2\u0158\u0156\3\2\2\2\u0158"+
		"\u0159\3\2\2\2\u0159\u015a\3\2\2\2\u015a\u015b\b\67\2\2\u015bn\3\2\2\2"+
		"\u015c\u015d\7^\2\2\u015d\u015e\t\t\2\2\u015ep\3\2\2\2\f\2\u00d8\u00e5"+
		"\u00ec\u00f1\u00f7\u0101\u010f\u0115\u0158\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}